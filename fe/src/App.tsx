import "./App.css";
import { Routes, Route } from "react-router-dom";
import TopPage from "./pages/TopPage";
import TaskManagementPage from "./pages/TaskManagementPage";
import LogPage from "./pages/LogPage";
import { useEffect, useRef, useState } from "react";

const App: React.FC = () => {
  const [message, setMessage] = useState<string>();
  const socketRef = useRef<WebSocket>();

  useEffect(() => {
    // #1.WebSocketオブジェクトを生成しサーバとの接続を開始
    const websocket = new WebSocket("ws://be-golang:8011/ws");
    socketRef.current = websocket;
    console.log("WebSocket connection initiated.");

    // 接続が成功したときのイベント
    websocket.onopen = () => {
      console.log("WebSocket connection established.");
    };

    // 接続エラーが発生したときのイベント
    websocket.onerror = (error) => {
      console.error("WebSocket connection error:", error);
    };

    // 接続がクローズされたときのイベント
    websocket.onclose = (event) => {
      console.log("WebSocket connection closed:", event);
    };

    // #2.メッセージ受信時のイベントハンドラを設定
    const onMessage = (event: MessageEvent<string>) => {
      console.log("Message received from server:", event.data);
      setMessage(event.data);
    };
    websocket.addEventListener("message", onMessage);

    // #3.useEffectのクリーンアップの中で、WebSocketのクローズ処理を実行
    return () => {
      websocket.close();
      websocket.removeEventListener("message", onMessage);
      console.log("WebSocket connection cleaned up.");
    };
  }, []);

  return (
    <div className="App">
      <Routes>
        <Route path="/" element={<TopPage />} />
        <Route path="/tm" element={<TaskManagementPage />} />
        <Route path="/log" element={<LogPage />} />
      </Routes>
    </div>
  );
};

export default App;
