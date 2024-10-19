import { TaskListPartProps } from "../types/componentsPropsTypes";
import TaskInfo from "./TaskInfo";
import "../styles/components/taskListPart.css";

const TaskListPart: React.FC<TaskListPartProps> = ({ statusName, tasks }) => {
  return (
    <div className="task-list-part">
      <p className="task-list-part-title">{statusName}</p>
      <div className="scroll-area">
        {tasks.map((task, index) => {
          return <TaskInfo task={task} key={index} />;
        })}
      </div>
    </div>
  );
};

export default TaskListPart;
