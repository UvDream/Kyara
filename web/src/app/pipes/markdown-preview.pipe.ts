import { Pipe, PipeTransform } from '@angular/core';
import * as  hljs from 'highlight.js';
import MarkdownIt from 'markdown-it';
@Pipe({
  name: 'MarkdownPreview'
})
export class MarkdownPreviewPipe implements PipeTransform {

  transform(value: any, args?: any): any {

    const md = new MarkdownIt({
      highlight: (str, lang) => {
        if (lang === '') {
          lang = 'javascript';
        }
        if (lang && hljs.getLanguage(lang)) {
          try {
            // 得到经过highlight.js之后的html代码
            const preCode = hljs.highlight(lang, str, true).value;
            // 以换行进行分割
            const lines = preCode.split(/\n/).slice(0, -1);
            // 添加自定义行号
            let html = lines.map((item, index) => {
              return '<li><span class="line-num" data-line="' + (index + 1) + '"></span>' + item + '</li>';
            }).join('');
            html = '<ol>' + html + '</ol>';
            // 添加代码语言
            if (lines.length) {
              html += '<div class="pBtn" ><b class="language">' + lang + '</b><b class="copyBtn">复制代码<div class="none"><textarea>' + str + '</textarea></div></b></div>';
            }

            return '<pre class="hljs"><code>' +
              html +
              '</code></pre>';
          } catch (__) { }
        }
      }
    }
    );
    // const renderer = new marked.Renderer();
    // renderer.heading = (text, level) => {
    //   const escapedText = text.toLowerCase().replace(/[^\w]+/g, '-');
    //   return `
    //             <h${level}>
    //               <a name="${escapedText}" class="anchor" href="#${escapedText}">
    //                 <span class="header-link"></span>
    //               </a>
    //               ${text}
    //             </h${level}>`;
    // };
    // marked.setOptions({
    //   renderer,
    //   highlight: (code, lang) => {
    //     console.log('高亮', lang);
    //     return hljs.highlightAuto(code).value;
    //   },
    //   pedantic: false,
    //   gfm: true,
    //   breaks: false,
    //   sanitize: false,
    //   smartLists: true,
    //   smartypants: false,
    //   xhtml: false
    // });
    // if (value && value.length > 0) {
    //   return marked(value);
    // }
    if (value && value.length > 0) {
      return md.render(value);
    }
    return null;
  }


}
