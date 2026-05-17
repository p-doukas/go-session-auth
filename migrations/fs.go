// The `FS` variable is an embedded filesystem that includes all `.sql`
// files in this folder. This allows the application or migration tools
// like Goose to access the migration files programmatically without
// needing the files to exist on disk at runtime.
package migrations

import "embed"

//go:embed *.sql
var FS embed.FS
