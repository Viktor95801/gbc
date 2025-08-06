$outfile = "gbc.exe"

function MakeAll {
    Write-Host "Building $outfile..."
    go build -o $outfile
} 
function MakeClean {
    Write-Host "Cleaning..."
    Remove-Item $outfile
}

if ($args[0] -eq "all") {
    MakeAll
} elseif ($args[0] -eq "clean") {
    MakeClean
} else {
    MakeAll
}

Write-Host "Done."
