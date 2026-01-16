package resources_test

import (
	"io/fs"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/firebolt-db/mcp-server/pkg/helpers/mimetype"
	"github.com/firebolt-db/mcp-server/pkg/resources"
)

func TestDocs_ResourceTemplate(t *testing.T) {

	docs := resources.NewDocs(nil, "")
	template := docs.ResourceTemplate()

	assert.NotEmpty(t, template.Name)
	assert.NotEmpty(t, template.URITemplate)
	assert.NotEmpty(t, template.MIMEType)
	assert.NotEmpty(t, template.Description)
}

func TestDocs_Handler_No(t *testing.T) {

	mockFS := newMockDocsFS()
	docs := resources.NewDocs(mockFS, "test-proof")

	t.Run("fetch overview article", func(t *testing.T) {

		request := &mcp.ReadResourceRequest{
			Params: &mcp.ReadResourceParams{
				Meta: map[string]any{
					"article": resources.DocsArticleOverview,
				},
			},
		}

		result, err := docs.Handler(t.Context(), request)
		assert.NoError(t, err)
		require.Len(t, result.Contents, 1)
		resource := result.Contents[0]
		assert.Equal(t, resources.DocsURI(resources.DocsArticleOverview), resource.URI)
		assert.Equal(t, mimetype.Markdown, resource.MIMEType)
		assert.Contains(t, resource.Text, "Foundational Knowledge Layer", "should contain the Foundational Knowledge Layer")
	})

	t.Run("fetch core overview article", func(t *testing.T) {

		request := &mcp.ReadResourceRequest{
			Params: &mcp.ReadResourceParams{
				Meta: map[string]any{
					"article": resources.DocsCoreArticleOverview,
				},
			},
		}

		result, err := docs.Handler(t.Context(), request)
		assert.NoError(t, err)
		require.Len(t, result.Contents, 1)
		resource := result.Contents[0]
		assert.Equal(t, resources.DocsURI(resources.DocsCoreArticleOverview), resource.URI)
		assert.Equal(t, mimetype.Markdown, resource.MIMEType)
		assert.Contains(t, resource.Text, "Foundational Knowledge Layer", "should contain the Foundational Knowledge Layer")
	})

	t.Run("fetch proof article", func(t *testing.T) {

		request := &mcp.ReadResourceRequest{
			Params: &mcp.ReadResourceParams{
				Meta: map[string]any{
					"article": resources.DocsArticleProof,
				},
			},
		}

		result, err := docs.Handler(t.Context(), request)
		assert.NoError(t, err)
		require.Len(t, result.Contents, 1)
		resource := result.Contents[0]
		assert.Equal(t, resources.DocsURI(resources.DocsArticleProof), resource.URI)
		assert.Equal(t, mimetype.Markdown, resource.MIMEType)
		assert.Contains(t, resource.Text, "test-proof", "should contain the proof value")
	})

	t.Run("fetch reference article", func(t *testing.T) {

		mockFS.On("ReadDir", ".").Return([]fs.DirEntry{
			newMockDirEntry("file1.md", false),
			newMockDirEntry("file2.md", false),
		}, nil)

		request := &mcp.ReadResourceRequest{
			Params: &mcp.ReadResourceParams{
				Meta: map[string]any{
					"article": resources.DocsArticleReference,
				},
			},
		}

		result, err := docs.Handler(t.Context(), request)
		assert.NoError(t, err)
		require.Len(t, result.Contents, 1)
		resource := result.Contents[0]
		assert.Equal(t, resources.DocsURI(resources.DocsArticleReference), resource.URI)
		assert.Equal(t, mimetype.Markdown, resource.MIMEType)
		assert.Contains(t, resource.Text, "file1.md")
		assert.Contains(t, resource.Text, "file2.md")
	})

	t.Run("fetch file-based article", func(t *testing.T) {

		mockFS.On("ReadFile", "file1.md").Return([]byte("# File 1 Content"), nil)

		request := &mcp.ReadResourceRequest{
			Params: &mcp.ReadResourceParams{
				Meta: map[string]any{
					"article": "file1.md",
				},
			},
		}

		result, err := docs.Handler(t.Context(), request)
		assert.NoError(t, err)
		require.Len(t, result.Contents, 1)
		resource := result.Contents[0]
		assert.Equal(t, resources.DocsURI("file1.md"), resource.URI)
		assert.Equal(t, mimetype.Markdown, resource.MIMEType)
		assert.Equal(t, "# File 1 Content", resource.Text)
	})
}

func newMockDocsFS() *mockDocsFS {
	return &mockDocsFS{}
}

type mockDocsFS struct {
	mock.Mock
}

func (m *mockDocsFS) Open(name string) (fs.File, error) {
	args := m.Called(name)
	return args.Get(0).(fs.File), args.Error(1)
}

func (m *mockDocsFS) ReadDir(name string) ([]fs.DirEntry, error) {
	args := m.Called(name)
	return args.Get(0).([]fs.DirEntry), args.Error(1)
}

func (m *mockDocsFS) ReadFile(name string) ([]byte, error) {
	args := m.Called(name)
	return args.Get(0).([]byte), args.Error(1)
}

func newMockDirEntry(name string, isDir bool) *mockDirEntry {
	return &mockDirEntry{name: name, isDir: isDir}
}

type mockDirEntry struct {
	name  string
	isDir bool
}

func (m *mockDirEntry) Name() string               { return m.name }
func (m *mockDirEntry) IsDir() bool                { return m.isDir }
func (m *mockDirEntry) Type() fs.FileMode          { return 0 }
func (m *mockDirEntry) Info() (fs.FileInfo, error) { return nil, nil }
