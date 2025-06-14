package fs

import "fmt"
// import "os/exec"

// HandleIntent handles file-system related intents.
func HandleIntent(name string, params map[string]string) (string, error) {
    // switch name {
    // case "list_files":
    //     dir := params["directory"]
    //     if dir == "" {
    //         dir = "." // Default to current directory if not specified
    //     }
    //     cmd := exec.Command("ls", dir)
    //     output, err := cmd.Output()
    //     if err != nil {
    //         return "", fmt.Errorf("error listing files in %s: %w", dir, err)
    //     }
    //     return string(output), nil
    // default:
    //     return "", fmt.Errorf("unsupported fs intent: %s", name)
    // }
    // fmt.Println("Handling intent:", name)
    // fmt.Println("With params:", params)
    
    return fmt.Sprintf("Handled fs intent: %s with params: %v", name, params), nil
}
