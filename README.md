Из-за ограничений по времени генерация OpenMetrics пока реализованно "примитивным" методом. Почти готов вывод валидного OpenMetrics с помощью библиотеки, но он требует небольшой конфигурации. Некоторые идеи по улучшению кода записаны в `TODO.md`.

# go-yaml2openmetrics

## Usage

There are 2 ways to use this service:
1. Run it as local server.
For this method you should have golang installed on your system.
a. Preparing for the build
`make init`
b. Run the code.
`make run`
c. If you want to stop container, press `Ctrl-C` in your terminal.

2. Run it as docker container.
a. Preparing for the build
`make init`
b. Build a container.
`make image`
c. Run the container (it will run interactively in your terminal).
`make start`
d. If you want to stop container, press `Ctrl-C` in your terminal.
e. If you want to delete container from your system after testing, please run:
`make delimage`

## Configuration

Configure `configs.yml` file to use custom host and .yml file name.