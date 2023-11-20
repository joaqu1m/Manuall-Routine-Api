import axios from "axios";

export default axios.create({
  baseURL: `${process.env.API_URL}/crm`,
  headers: {
    RoutineAuth: process.env.ROUTINE_AUTH,
  },
});
