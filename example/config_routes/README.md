# config_routes
Mock files in filesystem by yaml configuration

## usage
paste in `config.yaml` your api route and file for response
```yaml
/api/some/route:
  file: testdata/someresponse.json
/api/some/route/soap:
  file: testdata/someresponse.xml
```
run `gomockserve .` or `gomockserve . --c=otherconfiglocation.yaml`