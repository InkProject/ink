import { ACTION } from './index';
import * as menuAction from './editor';

export function saveContent(id, name, content) {
    return dispatch => {
        dispatch(menuAction.showLoading(true));
        fetch(`${apiURL}/articles/${id}`, {
            method: 'PUT',
            body: {
                name,
                content
            }
        }).then(function(response) {
            return response;
        }).then(function(data) {
            alert(data);
            dispatch(menuAction.showLoading(false));
        }).catch(function(error) {
            alert(JSON.stringify(error));
        });
    };
}
