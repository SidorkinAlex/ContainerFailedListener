# CyclicCommandCheckerAndExecutive

## Description

### English

**CyclicCommandCheckerAndExecutive** is a Go-based application designed to cyclically execute commands and manage their execution based on predefined conditions. It offers functionality to start, stop, restart, and create jobs through command-line interface parameters. The application can run as a daemon, ensuring continuous operation and monitoring.

#### Key Features
- **Cyclic Command Execution**: The application repeatedly runs commands specified in a configuration file.
- **Branching Logic**: Supports conditional execution of commands based on the results of previously executed commands.
- **Daemon Mode**: Operates as a background service.
- **Configurable**: Uses a JSON configuration file to define commands and execution intervals.
- **CLI Parameters**: Provides command-line options to start, stop, restart, or create jobs.
- **Logging**: Logs execution results and errors for debugging purposes.

#### Components
- **CliParamsParser**: Parses command-line arguments to determine the action to be performed.
- **Actions**: Contains functions for starting, stopping, restarting, and creating jobs.
- **Config**: Handles reading and parsing the JSON configuration file.
- **RunController**: Manages the daemon process and handles the stop signal.
- **FileUtils**: Utility functions for file operations.

#### Usage
1. **Configuration**: Define your commands in `/etc/CyclicCommandCheckerAndExecutive/config.json`.
2. **Start the Application**: Use `-start` parameter to start the cyclic execution.
3. **Stop the Application**: Use `-stop` parameter to stop the application.
4. **Restart the Application**: Use `-restart` parameter to restart the application.
5. **Create a Job**: Use `-create-job` parameter to create new job entries.

### Русский

**CyclicCommandCheckerAndExecutive** — это приложение на языке Go, предназначенное для циклического выполнения команд и управления их выполнением на основе заранее заданных условий. Оно предоставляет функционал для запуска, остановки, перезапуска и создания заданий через параметры командной строки. Программа может работать в режиме демона, обеспечивая непрерывную работу и мониторинг.

#### Основные возможности
- **Циклическое выполнение команд**: Программа повторно выполняет команды, указанные в конфигурационном файле.
- **Логика ветвления**: Поддерживает условное выполнение команд на основе результатов ранее выполненных команд.
- **Режим демона**: Работает как фоновый сервис.
- **Настраиваемость**: Использует конфигурационный файл JSON для определения команд и интервалов выполнения.
- **Параметры командной строки**: Предоставляет опции командной строки для запуска, остановки, перезапуска или создания заданий.
- **Логирование**: Ведет журнал результатов выполнения и ошибок для целей отладки.

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