import { LitElement, css, html } from 'lit'
import { customElement, property } from 'lit/decorators.js'

/**
 * サイトのボタンスタイル
 * 
 * @attribute outline アウトラインついたボタンにします。 
 */
@customElement('sau-button')
export class SauButton extends LitElement {
  @property({ attribute: 'outline', type: Boolean }) outline = false;
  @property({ attribute: 'message', type: String }) message = '';

  attributeChangedCallback() {
    this.outline = this.getAttribute("outline") === ''
    this.message = this.getAttribute("message") || ''
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

  static styles = css/* css */`
    /* reset css */
    button {
      background-color: transparent;
      border: none;
      cursor: pointer;
      outline: none;
      padding: 0;
      appearance: none;
    }

    .sau-button {
      display: flex;
      align-items: center;
      justify-content: center;
      width: 100%;
      max-width: 320px;
      height: 48px;
      padding: 8px 24px;
      font-size: 16px;
      color: #fff;
      text-align: center;
      overflow-wrap: anywhere;
      background-color: #6fa24a;
      border-radius: 32px;
      font-weight: bold;
    }

    @media (any-hover: hover) {
      .sau-button {
        transition: opacity 0.2s;
      }

      .sau-button:hover {
        opacity: 0.8;
      }
    }

    .sau-button.outline {
      background-color: white;
      color: #6fa24a;
      border: 1px solid #6fa24a;
    }
  `;
}

declare global {
  interface HTMLElementTagNameMap {
    'sau-button': HTMLButtonElement
  }
}