package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterDocsRoutes sets up Swagger UI and the raw OpenAPI YAML endpoint.
//
//	GET /swagger        → Swagger UI (HTML, served from CDN)
//	GET /swagger/doc.yaml → raw OpenAPI 3.0 YAML file
func RegisterDocsRoutes(server *gin.Engine) {
	// Serve the raw swagger.yaml so Swagger UI (and Backstage) can fetch it
	server.StaticFile("/swagger/doc.yaml", "./docs/swagger.yaml")

	// Serve Swagger UI for every sub-path so that the CDN assets load correctly
	server.GET("/swagger", swaggerUIHandler)
}

const swaggerUIHTML = `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Go Events REST API – Swagger UI</title>
    <link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@5/swagger-ui.css" />
  </head>
  <body>
    <div id="swagger-ui"></div>
    <script src="https://unpkg.com/swagger-ui-dist@5/swagger-ui-bundle.js"></script>
    <script>
      SwaggerUIBundle({
        url: "/swagger/doc.yaml",
        dom_id: "#swagger-ui",
        presets: [SwaggerUIBundle.presets.apis, SwaggerUIBundle.SwaggerUIStandalonePreset],
        layout: "BaseLayout",
        deepLinking: true,
        tryItOutEnabled: true,
      });
    </script>
  </body>
</html>`

func swaggerUIHandler(c *gin.Context) {
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(http.StatusOK, swaggerUIHTML)
}
