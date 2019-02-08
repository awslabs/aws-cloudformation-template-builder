# To do

* ~~Use go generate to build models for each resource type~~
    * Create `init()` func that registers the CFType funcs in `spec/types` to a map of CFType names to their respective funcs ie `map[name string]resourceFunc` and `map[name string]propertyFunc`
    * Add `go generate` tag to the main command
    * integrate `resourceFunc` and `propertyFunc` maps into main cfn-skeleton command
* Build dependencies automatically
* Interactive mode
