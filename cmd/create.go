/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create [name]",
	Short: "Create a new Go REST API project",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		fmt.Println("📁 Creating Go REST API project:", projectName)

		// Create directory structure
		dirs := []string{
			"cmd",
			"internal/handler",
			"internal/service",
			"internal/repository",
			"internal/domain",
		}

		for _, dir := range dirs {
			fullPath := filepath.Join(projectName, dir)
			if err := os.MkdirAll(fullPath, 0755); err != nil {
				fmt.Println("❌ Failed to create directory:", fullPath)
				return
			}
		}

		// Render main.go template
		mainTemplate := `package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.Run(":8080")
}`

		mainPath := filepath.Join(projectName, "cmd", "main.go")
		if err := os.WriteFile(mainPath, []byte(mainTemplate), 0644); err != nil {
			fmt.Println("❌ Failed to write main.go:", err)
			return
		}

		// Write go.mod
		goModContent := fmt.Sprintf("module %s\n\ngo 1.20\n\nrequire github.com/gin-gonic/gin v1.10.0\n", projectName)
		modPath := filepath.Join(projectName, "go.mod")
		if err := os.WriteFile(modPath, []byte(goModContent), 0644); err != nil {
			fmt.Println("❌ Failed to write go.mod:", err)
			return
		}

		fmt.Println("✅ Project created successfully.")
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
