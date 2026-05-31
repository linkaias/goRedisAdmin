import Vue from 'vue'
import VueI18n from 'vue-i18n'
import ElementLocale from 'element-ui/lib/locale'
import zhCNElement from 'element-ui/lib/locale/lang/zh-CN'
import enElement from 'element-ui/lib/locale/lang/en'
import zhCN from './messages/zh-CN'
import enUS from './messages/en-US'

Vue.use(VueI18n)

export const LANG_ZH = 'zh-CN'
export const LANG_EN = 'en-US'

const STORAGE_KEY = 'goRedisAdmin.lang'
const supported = [LANG_ZH, LANG_EN]

const messages = {
  [LANG_ZH]: zhCN,
  [LANG_EN]: enUS
}

const elementLocales = {
  [LANG_ZH]: zhCNElement,
  [LANG_EN]: enElement
}

const savedLang = localStorage.getItem(STORAGE_KEY)
const initialLang = supported.includes(savedLang) ? savedLang : LANG_ZH

const i18n = new VueI18n({
  locale: initialLang,
  fallbackLocale: LANG_EN,
  messages,
  silentTranslationWarn: true
})

export function setLanguage(lang) {
  const nextLang = supported.includes(lang) ? lang : LANG_ZH
  i18n.locale = nextLang
  ElementLocale.use(elementLocales[nextLang])
  localStorage.setItem(STORAGE_KEY, nextLang)
  return nextLang
}

export function getLanguage() {
  return i18n.locale
}

setLanguage(initialLang)

export default i18n
