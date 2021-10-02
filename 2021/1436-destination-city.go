func destCity(paths [][]string) string {
    g := make(map[string]string, len(paths))
    for _, p := range paths {
        g[p[0]] = p[1]
    }

    for _, p := range paths {
        if g[p[1]] == "" {
            return p[1]
        }
    }

    return ""
}
