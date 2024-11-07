import { useEffect, useState } from "react";
import config from "../config";
import axios from "axios";
import { TaskNoAt } from "../types/dataTypes";
import TaskListPart from "./TaskListPart";
import "../styles/components/taskList.css";

const TaskList = () => {
  const [pendingTasks, setPendingTasks] = useState<TaskNoAt[]>([]);
  const [inProgressTasks, setInProgressTasks] = useState<TaskNoAt[]>([]);
  const [doneTasks, setDoneTasks] = useState<TaskNoAt[]>([]);

  useEffect(() => {
    // タスク一覧の取得
    const getTasks = () => {
      const url: string = `${config.BACKEND_URL}/task`;
      axios
        .get(url)
        .then((response) => {
          const tasks: TaskNoAt[] = response.data.map((task: TaskNoAt) => ({
            ID: task.ID,
            Title: task.Title,
            Description: task.Description,
            StatusID: task.StatusID,
            Status: {
              ID: task.Status.ID,
              StatusName: task.Status.StatusName,
            },
            AddressID: task.AddressID,
            Address: {
              ID: task.Address.ID,
              IPAddress: task.Address.IPAddress,
              ColorCode: task.Address.ColorCode,
            },
          }));

          setPendingTasks(
            tasks.filter((task) => task.Status.StatusName === "Pending")
          );
          setInProgressTasks(
            tasks.filter((task) => task.Status.StatusName === "InProgress")
          );
          setDoneTasks(
            tasks.filter((task) => task.Status.StatusName === "Done")
          );
        })
        .catch((error) => {
          console.log(error);
          setPendingTasks([]);
          setInProgressTasks([]);
          setDoneTasks([]);
        });
    };

    getTasks();
  }, []);

  useEffect(() => {}, [pendingTasks, inProgressTasks, doneTasks]);

  return (
    <>
      <div className="flex-box">
        <TaskListPart statusName="保留中" tasks={pendingTasks} />
        <TaskListPart statusName="進行中" tasks={inProgressTasks} />
        <TaskListPart statusName="終了" tasks={doneTasks} />
      </div>
    </>
  );
};

export default TaskList;
