import { TaskInfoProps } from "../types/componentsPropsTypes";
import "../styles/components/taskInfo.css";

const TaskInfo: React.FC<TaskInfoProps> = ({ task }) => {
  return (
    <div className="task-info-div">
      <div
        className="circle"
        style={{ backgroundColor: task.Address.ColorCode }}
      ></div>
      <p className="task-title">
        {task.Title} <label>{task.Address.ColorCode}</label>
      </p>
      <p className="task-description">{task.Description}</p>
    </div>
  );
};

export default TaskInfo;
