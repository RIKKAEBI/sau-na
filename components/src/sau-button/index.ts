import { LitElement, html } from 'lit'
import { customElement, property } from 'lit/decorators.js'

/**
 * サイトのボタンスタイル
 * 
 * @attribute outline アウトラインついたボタンにします。 
 */
@customElement('sau-button')
export class SauButton extends LitElement {
  @property({ attribute: 'outline', type: Boolean }) outline = false;
  @property({ attribute: 'message', type: String }) message = ""
  
  protected createRenderRoot() {
    return this; // Light DOMを使うように
    // thisがLight DOMの<lit-component>のエレメントに当たります
  }

  attributeChangedCallback() {
    this.outline = !(this.getAttribute("outline") !== '')
    this.message = this.getAttribute("message") || ""
  }

  private _onClick() {
    console.log('hello world');
  }

  render() {
    const { outline, message } = this;

    return html/* html */`
      <button class=${(outline ? 'outline ' : '') + "sau-button"} @click=${this._onClick}>
        ${message}
      </button>
    `
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'sau-button': HTMLButtonElement
  }
}