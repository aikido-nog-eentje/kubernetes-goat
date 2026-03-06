package main

import (
	"fmt"
	"os/exec"
	"regexp"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html"
)

func main() {
	// Initialize HTML template engine
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Add panic recovery middleware
	app.Use(recover.New())

	// GET endpoint - render the main page
	app.Get("/", func(c *fiber.Ctx) error {
		// Rendering the templates
		return c.Render("index", fiber.Map{
			"Title": "Ping Your Servers",
		})
	})

	// POST endpoint - handle ping requests
	app.Post("/", func(c *fiber.Ctx) error {
		// Get the endpoint from form data
		endpoint := c.FormValue("endpoint")
		
		// Validate endpoint to prevent command injection
		// Only allow alphanumeric characters, dots, hyphens, and colons
		validEndpoint := regexp.MustCompile(`^[a-zA-Z0-9\.\-:]+$`)
		if !validEndpoint.MatchString(endpoint) {
			return fmt.Errorf("invalid input")
		}

		// Execute ping command with validated endpoint
		cmd := exec.Command("sh", "-c", "ping -c 2 "+endpoint)

		type error interface {
			Error() string
		}
		var finalOutput string
		// Run the command and capture output or error
		if output, err := cmd.Output(); err != nil {
			finalOutput = err.Error()
		} else {
			finalOutput = string(output)
		}

		// Render the result page with output
		return c.Render("index", fiber.Map{
			"Title":  "Ping Your Servers",
			"Output": finalOutput,
		})
	})

	// Start the server on port 80
	app.Listen("0.0.0.0:80")
}
