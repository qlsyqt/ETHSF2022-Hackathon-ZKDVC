import axios from "axios";
import { baseURL } from "../config";
import { message } from "antd";

const api = axios.create({
  baseURL,
});

api.interceptors.response.use((res) => {
  console.log("res");
  if (res.data.code === 200) {
    return res.data;
  } else {
    message.error(res.data.message);
    return null;
  }
});

const template = {
  list: async (params?: any) => api.get("/template", { params }),
  create: async (params?: any) => api.post("/template", params),
  getTem: async (id?: any) => api.get(`/template/${id}`),
  delete: async (id?: any) => api.delete(`/template/${id}`),
};

const offer = {
  list: async (params?: any) => api.get("/offer", { params }),
  create: async (params?: any) => api.post("/offer", params),
  delete: async (id?: any) => api.delete(`/offer/${id}`),
  patch: async (id?: any, params?: any) => api.patch(`/offer/${id}`, params),
  getQrCode: async (params?: any) => api.post(`/auth/create`, params),
};

export default {
  template,
  offer,
};
