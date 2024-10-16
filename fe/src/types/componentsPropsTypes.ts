import { TaskNoAt } from "./dataTypes";

export interface TaskListPartProps {
  statusName: string;
  tasks: TaskNoAt[];
}

export interface TaskListProps {
  statusName: string;
}

export interface TaskInfoProps {
  task: TaskNoAt;
}
