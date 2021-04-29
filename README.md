Из-за ограничений по времени генерация OpenMetrics пока реализованно "примитивным" методом. Почти готов вывод валидного OpenMetrics с помощью библиотеки, но он требует небольшой конфигурации. Некоторые идеи по улучшению кода записаны в `TODO.md`.

# go-yaml2openmetrics

## Usage

There are 2 ways to use this service:
### Run it as local server.

For this method you should have golang installed on your system.

1. Preparing for the build

```make init```

2. Run the code.

```make run```

3. If you want to stop container, press `Ctrl-C` in your terminal.


### Run it as docker container.
 
1. Preparing for the build

```make init```

2. Build a container.

```make image```

3. Run the container (it will run interactively in your terminal).

```make start```

4. If you want to stop container, press `Ctrl-C` in your terminal.

5. If you want to delete container from your system after testing, please run:

```make delimage```

## Configuration

Configure `configs.yml` file to use custom host and .yml file name.
