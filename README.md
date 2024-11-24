# CyclicCommandCheckerAndExecutive

## Description

**CyclicCommandCheckerAndExecutive** is a Go-based application designed to cyclically execute commands and manage their execution based on predefined conditions. It offers functionality to start, stop, restart, and create jobs through command-line interface parameters. The application can run as a daemon, ensuring continuous operation and monitoring.

#### Key Features
- **Cyclic Command Execution**: The application repeatedly runs commands specified in a configuration file.
- **Branching Logic**: Supports conditional execution of commands based on the results of previously executed commands.
- **Daemon Mode**: Operates as a background service.
- **Configurable**: Uses a JSON configuration file to define commands and execution intervals.
- **CLI Parameters**: Provides command-line options to start, stop, restart, or create jobs.
- **Logging**: Logs execution results and errors for debugging purposes.

## Setting up a command
to create a configuration file, use
```bash
nano /etc/CyclicCommandCheckerAndExecutive/config.json
```
Example of a configuration file:
```json
[
  {
    "checkingCommand": "cat /path/to/your/file.txt | tr -d '\n'",
    "interval": 5,
    "branchCommand": [
      {
        "resultExecution": "success",
        "commands": [
          "echo 'Command executed successfully.'",
          "echo '' > /path/to/your/file.txt"
        ]
      },
      {
        "resultExecution": "failure",
        "commands": [
          "bash command 1",
          "bash command 2"
        ]
      }
    ]
  }
]
```

In this example, cat is used to display the contents of the specified file /path/to/your/file.txt . You should replace /path/to/your/file.txt to the actual path to the file you want to track.
```bash
cat /path/to/your/file.txt
```
interval: Sets the time interval (in seconds) between each execution of the checkingCommand command. In this case, the value is 5, which means that the command will be executed every 5 seconds.

branchCommand command: This section defines the actions to be performed based on the result of the checkingCommand command.

resultExecution - The first object in the array indicates the actions to be performed if the string success (resultExecution: "success") was in the file. Here it includes:

commands - the second object is an array of commands that will be executed if the result of executing the checkingCommand command matches the expected resultExecution

In this example, Command executed successfully will be output to the application logs. and the file /path/to/your/file.txt it will be overwritten with an empty string. 

To use this configuration, save it in your configuration file (for example, /etc/CyclicCommandCheckerAndExecutive/config.json). The application will read this configuration and execute the cat command every 5 seconds, handling success and failure cases as defined.


#### Components
- **CliParamsParser**: Parses command-line arguments to determine the action to be performed.
- **Actions**: Contains functions for starting, stopping, restarting, and creating jobs.
- **Config**: Handles reading and parsing the JSON configuration file.
- **RunController**: Manages the daemon process and handles the stop signal.
- **FileUtils**: Utility functions for file operations.

#### Usage
#### sudo permissions are required to use the utility.

1. **Configuration**: Define your commands in `/etc/CyclicCommandCheckerAndExecutive/config.json`.
2. **Start the Application**: Use `-start` parameter to start the cyclic execution.
```bash
sudo ./build/CyclicCommandCheckerAndExecutive.0.0.1
```
3. **Stop the Application**: Use `-stop` parameter to stop the application.
4. **Restart the Application**: Use `-restart` parameter to restart the application.
5. **Create a Job**: Use `-create-job` parameter to create new job entries.

### Описание

**CyclicCommandCheckerAndExecutive** — это приложение на языке Go, предназначенное для циклического выполнения команд и управления их выполнением на основе заранее заданных условий. Оно предоставляет функционал для запуска, остановки, перезапуска и создания заданий через параметры командной строки. Программа может работать в режиме демона, обеспечивая непрерывную работу и мониторинг.

#### Основные возможности
#### для использования утилиты требуются разрешения sudo
- **Циклическое выполнение команд**: Программа повторно выполняет команды, указанные в конфигурационном файле.
- **Логика ветвления**: Поддерживает условное выполнение команд на основе результатов ранее выполненных команд.
- **Режим демона**: Работает как фоновый сервис.
- **Настраиваемость**: Использует конфигурационный файл JSON для определения команд и интервалов выполнения.
- **Параметры командной строки**: Предоставляет опции командной строки для запуска, остановки, перезапуска или создания заданий.
- **Логирование**: Ведет журнал результатов выполнения и ошибок для целей отладки.

## Настройка комманд
для создания конфигурационного файла используйте
```bash
nano /etc/CyclicCommandCheckerAndExecutive/config.json
```
Пример конфигурационного файла:
```json
[
  {
    "checkingCommand": "cat /path/to/your/file.txt | tr -d '\n'",
    "interval": 5,
    "branchCommand": [
      {
        "resultExecution": "success",
        "commands": [
          "echo 'Command executed successfully.'",
          "echo '' > /path/to/your/file.txt"
        ]
      },
      {
        "resultExecution": "failure",
        "commands": [
          "log_error",
          "send_alert"
        ]
      }
    ]
  }
]


```

В этом примере для отображения содержимого указанного файла используется cat /path/to/your/file.txt. Вам следует заменить /path/to/your/file.txt на фактический путь к файлу, который вы хотите отслеживать.
```bash
cat /path/to/your/file.txt
```
интервал: задает интервал времени (в секундах) между каждым выполнением команды checkingCommand. В данном случае значение равно 5, что означает, что команда будет выполняться каждые 5 секунд.

Команда branchCommand: В этом разделе определяются действия, которые необходимо выполнить на основе результата команды checkingCommand.

resultExecution - Первый объект в массиве указывает действия, которые необходимо выполнить в случае если в файле лежала строка success (resultExecution: "success"). Здесь он включает в себя:

commands - второй объект это массив команд, которые будут выполнены  в случае если результат выполнения команды checkingCommand будет совпадать с ожидаемым resultExecution

В данном примере будет выведено в логи приложения Command executed successfully. и файл /path/to/your/file.txt будет перезаписан пустой строкой. 

Чтобы использовать эту конфигурацию, сохраните ее в своем конфигурационном файле (например, /etc/CyclicCommandCheckerAndExecutive/config.json). Приложение будет считывать эту конфигурацию и выполнять команду cat каждые 5 секунд, обрабатывая случаи успешного выполнения и сбоя, как определено.


#### Компоненты
- **CliParamsParser**: Разбирает аргументы командной строки для определения действия, которое необходимо выполнить.
- **Actions**: Содержит функции для запуска, остановки, перезапуска и создания заданий.
- **Config**: Обрабатывает чтение и разбор конфигурационного файла JSON.
- **RunController**: Управляет процессом демона и обрабатывает сигнал остановки.
- **FileUtils**: Утилиты для работы с файлами.

#### Использование
1. **Конфигурация**: Определите свои команды в файле `/etc/CyclicCommandCheckerAndExecutive/config.json`.
2. **Запуск приложения**: Используйте параметр `-start` для начала циклического выполнения.
3. **Остановка приложения**: Используйте параметр `-stop` для остановки приложения.
4. **Перезапуск приложения**: Используйте параметр `-restart` для перезапуска приложения.
5. **Создание задания**: Используйте параметр `-create-job` для создания новых записей заданий.
