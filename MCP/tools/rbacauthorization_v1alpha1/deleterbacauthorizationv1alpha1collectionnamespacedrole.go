package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/kubernetes/mcp-server/config"
	"github.com/kubernetes/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Deleterbacauthorizationv1alpha1collectionnamespacedroleHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["continue"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("continue=%v", val))
		}
		if val, ok := args["fieldSelector"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fieldSelector=%v", val))
		}
		if val, ok := args["includeUninitialized"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("includeUninitialized=%v", val))
		}
		if val, ok := args["labelSelector"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("labelSelector=%v", val))
		}
		if val, ok := args["limit"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("limit=%v", val))
		}
		if val, ok := args["resourceVersion"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("resourceVersion=%v", val))
		}
		if val, ok := args["timeoutSeconds"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("timeoutSeconds=%v", val))
		}
		if val, ok := args["watch"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("watch=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/apis/rbac.authorization.k8s.io/v1alpha1/namespaces/%s/roles%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("DELETE", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// No authentication required for this endpoint
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result map[string]interface{}
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateDeleterbacauthorizationv1alpha1collectionnamespacedroleTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_apis_rbac.authorization.k8s.io_v1alpha1_namespaces_namespace_roles",
		mcp.WithDescription("delete collection of Role"),
		mcp.WithString("continue", mcp.Description("The continue option should be set when retrieving more results from the server. Since this value is server defined, clients may only use the continue value from a previous query result with identical query parameters (except for the value of continue) and the server may reject a continue value it does not recognize. If the specified continue value is no longer valid whether due to expiration (generally five to fifteen minutes) or a configuration change on the server the server will respond with a 410 ResourceExpired error indicating the client must restart their list without the continue field. This field is not supported when watch is true. Clients may start a watch from the last resourceVersion value returned by the server and not miss any modifications.")),
		mcp.WithString("fieldSelector", mcp.Description("A selector to restrict the list of returned objects by their fields. Defaults to everything.")),
		mcp.WithString("includeUninitialized", mcp.Description("If true, partially initialized resources are included in the response.")),
		mcp.WithString("labelSelector", mcp.Description("A selector to restrict the list of returned objects by their labels. Defaults to everything.")),
		mcp.WithString("limit", mcp.Description("limit is a maximum number of responses to return for a list call. If more items exist, the server will set the `continue` field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results. Setting a limit may return fewer than the requested amount of items (up to zero items) in the event all requested objects are filtered out and clients should only use the presence of the continue field to determine whether more results are available. Servers may choose not to support the limit argument and will return all of the available results. If limit is specified and the continue field is empty, clients may assume that no more results are available. This field is not supported if watch is true.\n\nThe server guarantees that the objects returned when using continue will be identical to issuing a single list call without a limit - that is, no objects created, modified, or deleted after the first request is issued will be included in any subsequent continued requests. This is sometimes referred to as a consistent snapshot, and ensures that a client that is using limit to receive smaller chunks of a very large result can ensure they see all possible objects. If objects are updated during a chunked list the version of the object that was present at the time the first list result was calculated is returned.")),
		mcp.WithString("resourceVersion", mcp.Description("When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history. When specified for list: - if unset, then the result is returned from remote storage based on quorum-read flag; - if it's 0, then we simply return what we currently have in cache, no guarantee; - if set to non zero, then the result is at least as fresh as given rv.")),
		mcp.WithString("timeoutSeconds", mcp.Description("Timeout for the list/watch call. This limits the duration of the call, regardless of any activity or inactivity.")),
		mcp.WithString("watch", mcp.Description("Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Deleterbacauthorizationv1alpha1collectionnamespacedroleHandler(cfg),
	}
}
