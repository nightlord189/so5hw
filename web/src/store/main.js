/* eslint-disable max-len */
/* eslint-disable no-else-return */
/* eslint-disable no-mixed-operators */
import { makeAutoObservable } from 'mobx';

class MainStore {
    constructor() {
        makeAutoObservable(this);
    }
}

export default new MainStore();