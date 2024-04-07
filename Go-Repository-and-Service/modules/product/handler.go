package product

import (
	"eduwork-bimo/Go-Repository-and-Service/helper"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service Service
}

func NewProductHandler(service Service) *Handler {
	return &Handler{service}
}

func (h *Handler) CreateProduct(c *fiber.Ctx) error {
	var product Product
	if err := c.BodyParser(&product); err != nil {
		response := helper.APIResponse("Failed to parse request body", fiber.StatusBadRequest, "error", nil)
		return c.Status(response.Meta.Code).JSON(response)
	}

	if err := h.service.CreateProduct(&product); err != nil {
		response := helper.APIResponse("Failed to create product", fiber.StatusInternalServerError, "error", nil)
		return c.Status(response.Meta.Code).JSON(response)
	}

	response := helper.APIResponse("Product create successfully", fiber.StatusOK, "success", product)
	return c.JSON(response)
}

func (h *Handler) GetAllProducts(c *fiber.Ctx) error {
	products, err := h.service.GetAllProducts()
	if err != nil {
		response := helper.APIResponse("Failed to fetch products", fiber.StatusInternalServerError, "error", nil)
		return c.Status(response.Meta.Code).JSON(response)
	}

	response := helper.APIResponse("Products fetched succesfully", fiber.StatusOK, "success", products)
	return c.JSON(response)
}

func (h *Handler) GetProductById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		response := helper.APIResponse("Invalid product ID", fiber.StatusBadRequest, "error", nil)
		return c.Status(response.Meta.Code).JSON(response)
	}

	product, err := h.service.GetProductById(id)
	if err != nil {
		response := helper.APIResponse("Failed to fetch product", fiber.StatusInternalServerError, "error", nil)
		return c.Status(response.Meta.Code).JSON(response)
	}

	response := helper.APIResponse("Products fetched succesfully", fiber.StatusOK, "success", product)
	return c.JSON(response)
}

func (h *Handler) UpdateProduct(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		response := helper.APIResponse("Invalid product ID", fiber.StatusBadRequest, "error", nil)
		return c.Status(response.Meta.Code).JSON(response)
	}

	var updatedProduct Product
	if err := c.BodyParser(&updatedProduct); err != nil {
		response := helper.APIResponse("Failed to parse request body", fiber.StatusBadRequest, "error", nil)
		return c.Status(response.Meta.Code).JSON(response)
	}

	updatedProduct.ID = id
	if err := h.service.UpdateProduct(&updatedProduct); err != nil {
		response := helper.APIResponse("Failed to update product", fiber.StatusInternalServerError, "error", nil)
		return c.Status(response.Meta.Code).JSON(response)
	}

	response := helper.APIResponse("Product updated succesfully", fiber.StatusOK, "success", updatedProduct)
	return c.JSON(response)
}

func (h *Handler) DeleteProduct(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		response := helper.APIResponse("Invalid product ID", fiber.StatusBadRequest, "error", nil)
		return c.Status(response.Meta.Code).JSON(response)
	}

	if err := h.service.DeleteProduct(id); err != nil {
		if err.Error() == "Product not found" {
			response := helper.APIResponse("Product not found", fiber.StatusNotFound, "error", nil)
			return c.Status(response.Meta.Code).JSON(response)
		}

		response := helper.APIResponse("Failed to delete product", fiber.StatusInternalServerError, "error", nil)
		return c.Status(response.Meta.Code).JSON(response)
	}

	response := helper.APIResponse("Product deleted succesfully", fiber.StatusOK, "success", nil)
	return c.JSON(response)
}
