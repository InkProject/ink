import React from 'react';
import Component from './index';
import ace from 'brace';
import 'brace/mode/markdown';
import 'brace/theme/tomorrow';

export default class Editor extends Component {
    resizeEditor () {
        let width = window.innerWidth ||
            document.documentElement.clientWidth ||
            document.body.clientWidth;
        if (width > 800) {
            this.editor.renderer.setPadding((width - 800) / 2);
        }
    }
    componentDidUpdate () {
        this.editor.setValue(this.props.content, -1);
    }
    componentDidMount () {
        let editor = ace.edit('editor');
        editor.setOptions({
            scrollPastEnd: true,
            showGutter: false,
            wrap: true,
            theme: 'ace/theme/tomorrow',
            mode: 'ace/mode/markdown',
            showPrintMargin: false,
            fontSize: '14px',
            fontFamily: "Menlo, Consolas, 'source-code-pro', 'DejaVu Sans Mono', Monaco, 'Ubuntu Mono', 'Courier New', Courier, 'Microsoft Yahei', 'Hiragino Sans GB', 'WenQuanYi Micro Hei', monospace",
            hScrollBarAlwaysVisible: false,
            selectionStyle: "line",
            displayIndentGuides: false
        });
        editor.renderer.setScrollMargin(200, 200);
        editor.container.style.lineHeight = 2;
        editor.on('input', () =>
            this.props.onChange(this.editor.getValue())
        );
        this.editor = editor;
        this.resizeEditor();
        window.addEventListener('resize', this.resizeEditor.bind(this));
    }
    render() {
        return (
            <div id="editor"></div>
        );
    }
}
