- [x] Using prometheus libraby to generate valid OpenMetrics with HandlerOpts{EnableOpenMetrics: true}
- [ ] Use buffered string to concatenate output in GenerateOpenMetricsText function.
- [ ] Parsing YAML file using map[string]interface{}, so we can parse __any type of YAML file__ with recursive algorithm. In this algo the value of map can be one of 3 different types: map, slice, openmetrics value. So we can use type switch and 3 different function to parse YAML tree.
  - [x] Parsing YAML file using map[interface{}]interface{}
  - [ ] Add support for different openmetrics value.
- [ ] Run http server in separate __go routine__. Generate openmetrics data and send to server using __buffered channel__.
- [ ] Create _Domain Specific Language_ for OpenMetrics types.
- [ ] Error management: 1) add errors return to functions where needed; 2) grouping your errors into a separate module; 3) Moving error messages to configuration
- [ ] Develop the Source abstraction layer to be able to use different types of yaml sources.
- [ ] Adding some more comments
- [ ] Testing
