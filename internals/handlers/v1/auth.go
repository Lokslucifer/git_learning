package v1

import (
	"errors"
	"large_fss/internals/constants"
	customerrors "large_fss/internals/customErrors"
	"large_fss/internals/dto"
	"large_fss/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SignupHandler(c *gin.Context) {

	var signupRequest dto.SignUpDTO

	err := c.BindJSON(&signupRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"message": customerrors.ErrBadRequest.Error(),
			},
		})
		return
	}
	userToken, err := h.ser.Signup(c, &signupRequest)

	if err != nil {
		if errors.Is(err, customerrors.ErrUserAlreadyExists) {
			c.JSON(http.StatusConflict, gin.H{
				"error": gin.H{
					"message": customerrors.ErrUserAlreadyExists.Error(),
				},
			})
			return
		} else {

			utils.LogErrorWithStack(c, "Internal Server Error in Signup user", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": gin.H{
					"message": customerrors.ErrInternalServer.Error(),
				},
			})

			return
		}
	}
	c.SetCookie(
		"auth_token", // Cookie name
		userToken,    // Value
		3600*6,       // MaxAge in seconds (e.g., 1 hour)
		"/",          // Path
		"",           // Domain ("" uses the request's domain)
		false,        // Secure (set to true in production with HTTPS)
		true,         // HttpOnly (not accessible via JS)
	)
	c.JSON(http.StatusOK, gin.H{
		"message": "Logout successful",
	})

}

func (h *Handler) LoginHandler(c *gin.Context) {

	var loginRequest dto.LoginDTO

	err := c.BindJSON(&loginRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"message": customerrors.ErrInvalidInput.Error(),
			},
		})
		return
	}

	userToken, err := h.ser.Login(c, &loginRequest)
	if err != nil {
		if errors.Is(err, customerrors.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": gin.H{
					"message": customerrors.ErrUserNotFound.Error(),
				},
			})
			return
		} else if errors.Is(err, customerrors.ErrInvalidPassword) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": gin.H{
					"message": customerrors.ErrInvalidPassword.Error(),
				},
			})
			return
		} else {
			utils.LogErrorWithStack(c, "Internal Server Error in Login", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": gin.H{
					"message": customerrors.ErrInternalServer.Error(),
				},
			})
			return
		}
	}
	c.SetCookie(
		"auth_token", // Cookie name
		userToken,    // Value
		3600*6,       // MaxAge in seconds (e.g., 1 hour)
		"/",          // Path
		"",           // Domain ("" uses the request's domain)
		false,        // Secure (set to true in production with HTTPS)
		true,         // HttpOnly (not accessible via JS)
	)

	c.JSON(http.StatusCreated, gin.H{
		"message": constants.SuccessMessage,
		"Token":   userToken,
	})

}

func (h *Handler) LogoutHandler(c *gin.Context) {
	c.SetCookie(
		"auth_token", // Cookie name
		"",           // Empty value
		-1,           // MaxAge < 0 deletes the cookie
		"/",          // Path must match the original
		"",           // Domain (empty means current domain)
		false,        // Secure (true in production)
		true,         // HttpOnly
	)

	c.JSON(http.StatusOK, gin.H{
		"message": "Logout successful",
	})
}
