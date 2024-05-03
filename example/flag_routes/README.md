# flag_routes
Mock files in filesystem by flags

temporary works only by <route>:<file> configuration
for more precise setup use yaml-config options
## usage
run 
`gomockserve . --r="/api/some:testdata/someresponse.json" --r="/api/soap:testdata/someresponse.xml"`
or with base path
`gomockserve ./testdata --r="/api/some:someresponse.json" --r="/api/soap:someresponse.xml"`
