import PageTitle from "../components/PageTitle";
import TaskList from "../components/TaskList";

const TaskManagementPage: React.FC = () => {
  return (
    <>
      <PageTitle title="タスク管理" />
      <TaskList />
    </>
  );
};

export default TaskManagementPage;
