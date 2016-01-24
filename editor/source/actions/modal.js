import { ACTION, utilAction, listAction } from './index';
import history from '../history';

export function createArticle(name) {
    return dispatch => {
        dispatch(utilAction.showTip('load'));
        utilAction.apiRequest('POST', `articles`, {
            name,
            content: ''
        }).then(function(data) {
            dispatch(utilAction.showTip('auto', '创建成功'));
            dispatch(listAction.fetchList());
            history.replaceState(null, `/edit/${data.id}`);
        });
    };
}
