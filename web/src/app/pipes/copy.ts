const _ = require('lodash');

try {
    // Node js will throw an error
    // tslint:disable-next-line:no-unused-expression
    this === window;
}
catch (_err) {
}

const defaultOptions = {
    iconStyle: 'font-size: 21px; opacity: 0.4;',
    iconClass: 'mdi mdi-content-copy',
    buttonStyle: 'position: absolute; top: 7.5px; right: 6px; cursor: pointer; outline: none;',
    buttonClass: ''
};

// tslint:disable-next-line:typedef
function renderCode(origRule, options) {
    options = _.merge(defaultOptions, options);
    return (...args) => {
        const [tokens, idx] = args;
        const content = tokens[idx].content;
        const origRendered = origRule(...args);
        if (content.length === 0) {
            return origRendered;
        }
        return `
        <div style="position: relative">
            ${origRendered}
            <button  class="markdown-it-code-copy" data-clipboard-text="1111"   title="Copy" >
                <span class="markdown-it-code-copy" data-clipboard-text="1111" style="${options.iconStyle}"  >拷贝</span>
            </button>
        </div>
`;
    };
}


module.exports = (md, options) => {
    md.renderer.rules.code_block = renderCode(md.renderer.rules.code_block, options);
    md.renderer.rules.fence = renderCode(md.renderer.rules.fence, options);
};