import {Configuration} from "./configuration.ts";
import {AuthApiFactory} from "./api.ts";

const cfg = new Configuration()

export const apiService = AuthApiFactory(cfg)