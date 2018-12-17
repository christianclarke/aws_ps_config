# aws_ps_config
This golang program will
- fetch data from AWS ParameterStore given a specific path and region.
- Set the data as environment vars whch are visible only within the context of the executing program

## Use cases.
The progrma could be used to inject ParamStore config/secrets into a running program or container
