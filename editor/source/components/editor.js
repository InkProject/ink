import React from 'react';

import ace from 'brace';
import 'brace/mode/markdown';
import 'brace/theme/tomorrow';

export default class Editor extends React.Component {
    componentWillReceiveProps (props) {
        this.editor.setValue(props.content, -1);
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
            displayIndentGuides: false
        });
        editor.renderer.setScrollMargin(200, 200);
        window.addEventListener('resize', function() {
            let width = window.innerWidth ||
                document.documentElement.clientWidth ||
                document.body.clientWidth;
            if (width > 700) {
                editor.renderer.setPadding((width - 700) / 2);
            }
        });
        editor.container.style.lineHeight = 2;
        this.editor = editor;
    }
    render() {
        return (
            <div id="editor"></div>
        );
    }
}
