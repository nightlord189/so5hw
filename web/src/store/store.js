import { atom } from 'jotai';
import { atomWithStorage } from 'jotai/utils'

const authTokenAtom = atomWithStorage('authToken', '');
const usernameAtom = atomWithStorage('username', '');
const userTypeAtom = atomWithStorage('userType', 'customer');

const categoriesAtom = atom([]);

export {authTokenAtom, usernameAtom, userTypeAtom, categoriesAtom};