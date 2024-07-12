# Replace these values with appropriate ones
param (
    [string]$NEW_MODULE_NAME = ""
)

$OLD_MODULE = "gin-api"

if ($NEW_MODULE_NAME -eq "") {
    Write-Host "Please enter the new module name:"
    $NEW_MODULE_NAME = Read-Host
    if ($NEW_MODULE_NAME -eq "") {
        Write-Host "No module name provided. Exiting script."
        exit 1
    }
}

go mod edit -module $NEW_MODULE_NAME

# Rename all imported modules in .go files
Get-ChildItem -Path . -Filter '*.go' -Recurse | ForEach-Object {
    $content = Get-Content -Path $_.FullName
    $updatedContent = $content -replace $OLD_MODULE, $NEW_MODULE_NAME
    $updatedContent | Set-Content -Path $filePath
}