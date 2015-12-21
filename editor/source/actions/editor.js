import { ACTION } from './index';
import yaml from 'js-yaml';

let parseConfig = function(data) {
    let dataAry = data.split('---');
    try {
        return yaml.safeLoad(dataAry[0]);
    } catch (err) {
        return null;
    }
};

export function setHeader(data) {
    let config = parseConfig(data);
    let { title, tags } = config || {};
    return {
        type: ACTION.SET_HEADER,
        title,
        tags
    }
}

export function setContent(content) {
    return {
        type: ACTION.SET_CONTENT,
        content
    }
}
