import { TaskListPartProps } from "../types/componentsPropsTypes";

const TaskListPart: React.FC<TaskListPartProps> = ({ statusName, tasks }) => {
  return (
    <div>
      <p>{statusName}</p>
      {tasks.map((task, index) => {
        return <div key={index}>{task.Address.ColorCode}</div>;
      })}
    </div>
  );
};

export default TaskListPart;
