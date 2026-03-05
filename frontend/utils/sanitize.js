const ALLOWED_TAGS = new Set([
  'p', 'br', 'b', 'i', 'em', 'strong', 'u', 's', 'del',
  'h1', 'h2', 'h3', 'h4', 'h5', 'h6',
  'ul', 'ol', 'li', 'blockquote', 'pre', 'code',
  'a', 'img', 'table', 'thead', 'tbody', 'tr', 'th', 'td',
  'div', 'span', 'hr', 'sup', 'sub', 'figure', 'figcaption', 'video', 'source'
])

const ALLOWED_ATTRS = new Set([
  'href', 'src', 'alt', 'title', 'width', 'height',
  'class', 'style', 'target', 'rel', 'colspan', 'rowspan',
  'controls', 'autoplay', 'type'
])

export function sanitizeHtml(html) {
  if (!html) return ''
  if (typeof window === 'undefined') return html

  const doc = new DOMParser().parseFromString(html, 'text/html')
  sanitizeNode(doc.body)
  return doc.body.innerHTML
}

function sanitizeNode(node) {
  const children = Array.from(node.childNodes)
  for (const child of children) {
    if (child.nodeType === 1) {
      const tag = child.tagName.toLowerCase()
      if (tag === 'script' || tag === 'iframe' || tag === 'object' || tag === 'embed' || !ALLOWED_TAGS.has(tag)) {
        child.remove()
        continue
      }
      const attrs = Array.from(child.attributes)
      for (const attr of attrs) {
        const name = attr.name.toLowerCase()
        if (!ALLOWED_ATTRS.has(name)) {
          child.removeAttribute(attr.name)
          continue
        }
        if ((name === 'href' || name === 'src') && attr.value.trim().toLowerCase().startsWith('javascript')) {
          child.removeAttribute(attr.name)
        }
        if (name === 'style') {
          child.setAttribute('style', attr.value.replace(/expression|javascript|vbscript/gi, ''))
        }
      }
      if (tag === 'a') {
        child.setAttribute('rel', 'noopener noreferrer')
        child.setAttribute('target', '_blank')
      }
      sanitizeNode(child)
    }
  }
}
