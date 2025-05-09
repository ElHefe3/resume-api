package utils

import "strings"

func GetFileType(filename string) string {
    switch {
    case strings.HasSuffix(filename, ".md"):
        return "markdown"
    case strings.HasSuffix(filename, ".pdf"):
        return "pdf"
    case strings.HasSuffix(filename, ".html"):
        return "html"
    default:
        return "unknown"
    }
}
