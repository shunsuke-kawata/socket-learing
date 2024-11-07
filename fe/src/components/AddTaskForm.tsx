import { useRef, useState } from "react";
import axios from "axios";
import { TaskParam } from "../types/dataTypes";
import "../styles/components/addTaskForm.css";
import config from "../config";

const AddTaskForm: React.FC = () => {
  const [isOpenForm, setIsOpenForm] = useState<boolean>(false);

  // 1つのuseRefにまとめてオブジェクト形式で参照
  const formRefs = useRef({
    taskTitle: "",
    taskDescription: "",
  });

  const handleCancelButton = () => {
    console.log("cancel add task");
    setIsOpenForm(false);
  };

  const handleSubmitButton = () => {
    console.log("submit add task", formRefs.current);
    // formRefs.current.taskTitle, formRefs.current.taskDescription を使って値を取得
    const param: TaskParam = {
      Title: formRefs.current.taskTitle,
      Description: formRefs.current.taskDescription,
    };
    const url: string = `${config.BACKEND_URL}/task`;
    axios
      .post(url, param)
      .then((response) => {
        console.log(response);
        formRefs.current.taskTitle = "";
        formRefs.current.taskDescription = "";
      })
      .catch((error) => console.log("error", error));
    setIsOpenForm(false);
  };

  return (
    <>
      <div className="add-task-form">
        <div
          className="show-form-button"
          onClick={() => setIsOpenForm(!isOpenForm)}
        >
          <span className={isOpenForm ? "dli-minus" : "dli-plus"}></span>
        </div>
        {isOpenForm ? (
          <div className="task-form-main">
            <label className="task-form-label">
              タスク名
              <input
                type="text"
                className="task-form-input"
                defaultValue=""
                onChange={(e) => (formRefs.current.taskTitle = e.target.value)}
              />
            </label>
            <label className="task-form-label">
              説明
              <input
                type="text"
                className="task-form-input"
                defaultValue=""
                onChange={(e) =>
                  (formRefs.current.taskDescription = e.target.value)
                }
              />
            </label>
            <div className="submit-button">
              <input
                type="button"
                className="submit-button-common"
                value="キャンセル"
                onClick={handleCancelButton}
              />
              <input
                type="button"
                className="submit-button-common"
                value="追加"
                onClick={handleSubmitButton}
              />
            </div>
          </div>
        ) : (
          <></>
        )}
      </div>
    </>
  );
};

export default AddTaskForm;
