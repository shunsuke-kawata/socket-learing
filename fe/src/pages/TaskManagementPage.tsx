import PageTitle from "../components/PageTitle";
import TaskList from "../components/TaskList";
import AddTaskForm from "../components/AddTaskForm";

const TaskManagementPage: React.FC = () => {
  return (
    <>
      <PageTitle title="タスク管理" />
      <TaskList />
      <AddTaskForm />
    </>
  );
};

export default TaskManagementPage;
