# TODO

## Task Tacker CLI

### 1. Tasks:

#### 1.1 Task Entity:

- [x] Create task entity;
  - [x] id;
  - [x] description;
  - [x] status;
  - [x] createdAt;
  - [x] updatedAt;
- [x] Create task status type;
- [ ] Add task props validation;
  - [ ] Description should not be empty;
- [ ] Create a way to calc the serial uint id;
- [ ] Add unit tests;

#### 1.2 TaskList Entity:

- [x] Create task list entity;
  - [x] tasks;
- [x] Create a get method;
- [x] Create a set method;
- [x] Make task list persistable;
- [ ] Doc the task list entity code;
- [ ] Add unit tests.

### 2. Data Source:

- [x] Create a data sourcer interface;
- [x] Create a data persister interface;
- [x] Create a data source of json kind;
- [ ] Add unit tests.

### 3. CLI:

- [ ] Make possible to create a task;
- [ ] Make possible to udate a task;
- [ ] Make possible to delete a task;
- [ ] Make possible to mark a task as in progress or done;
- [ ] Make possible to list all tasks;
- [ ] Make possible to list all todo tasks;
- [ ] Make possible to list all in progress tasks;
- [ ] Make possible to list all done tasks;
- [ ] Add cli doc (help) for the entry proint and each method;

### 4. Options:

- [ ] Make it possible to change de default path to the JSON database.

### 5. Support and Compile:

- [ ] Support and compile for Linux;
- [ ] Support and compile for Darwin;
- [ ] Support and compile for Windows;
