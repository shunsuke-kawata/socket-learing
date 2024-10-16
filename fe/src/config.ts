const config = {
  frontendUrl: process.env.REACT_APP_FRONTEND_URL,
  backendUrl: process.env.REACT_APP_BACKEND_URL,
  statusDict: {
    Pending: 1,
    InProgress: 2,
    Done: 3,
  },
};

export default config;
