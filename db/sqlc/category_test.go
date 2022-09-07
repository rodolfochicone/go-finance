package db

import (
	"context"
	"github.com/cip8/autoname"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomCategory(t *testing.T) Category {
	user := createRandomUser(t)
	arg := CreateCategoryParams{
		UserID:      user.ID,
		Title:       autoname.Generate(""),
		Type:        "debit",
		Description: autoname.Generate(" "),
	}

	category, err := testQueries.CreateCategory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, category)

	require.Equal(t, arg.UserID, category.UserID)
	require.Equal(t, arg.Title, category.Title)
	require.Equal(t, arg.Type, category.Type)
	require.Equal(t, arg.Description, category.Description)
	require.NotEmpty(t, category.CreatedAt)

	return category
}

func TestCreateCategory(t *testing.T) {
	createRandomCategory(t)
}

func TestGetCategoryById(t *testing.T) {
	category1 := createRandomCategory(t)

	category2, err := testQueries.GetCategoriesById(context.Background(), category1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, category2)

	require.Equal(t, category1.UserID, category2.UserID)
	require.Equal(t, category1.Title, category2.Title)
	require.Equal(t, category1.Type, category2.Type)
	require.Equal(t, category1.Description, category2.Description)
	require.NotEmpty(t, category2.CreatedAt)
}

func TestDeleteCategory(t *testing.T) {
	category1 := createRandomCategory(t)

	err := testQueries.DeleteCategory(context.Background(), category1.ID)
	require.NoError(t, err)

	category2, err := testQueries.GetCategoriesById(context.Background(), category1.ID)
	require.Error(t, err)
	require.Empty(t, category2)
}

func TestUpdateCategory(t *testing.T) {
	category1 := createRandomCategory(t)

	arg := UpdateCategoryParams{
		ID:          category1.ID,
		Title:       autoname.Generate(""),
		Description: autoname.Generate(" "),
	}

	category2, err := testQueries.UpdateCategory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, category2)

	require.Equal(t, category1.ID, category2.ID)
	require.Equal(t, arg.Title, category2.Title)
	require.Equal(t, arg.Description, category2.Description)
	require.NotEmpty(t, category2.CreatedAt)
}

func TestListCategories(t *testing.T) {
	lastCategory := createRandomCategory(t)

	arg := GetCategoriesByUserIdParams{
		UserID:      lastCategory.UserID,
		Type:        lastCategory.Type,
		Title:       lastCategory.Title,
		Description: lastCategory.Description,
	}

	categories, err := testQueries.GetCategoriesByUserId(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, categories, 1)
	require.NotEmpty(t, categories)

	for _, category := range categories {
		require.Equal(t, lastCategory.ID, category.ID)
		require.Equal(t, lastCategory.Title, category.Title)
		require.Equal(t, lastCategory.Description, category.Description)
		require.NotEmpty(t, lastCategory.CreatedAt)
	}
}
