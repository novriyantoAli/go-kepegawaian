import { handleResponse, requestOptions } from '@/_helpers';

export const usersService = {
    getAll
};

function getAll() {
    return fetch(`http://localhost:3000/api/pegawai`, requestOptions.get())
        .then(handleResponse);
}