//#Libro frontend, template CoderDocs
import Vue from 'vue'
import VueI18n from 'vue-i18n'

const messages = {
    en: {
        message: {
            wait: 'Please wait...',
            error_product_not_found: 'No product mapped to domain "{domain}".',
        }
    },
    vi: {
        message: {
            wait: 'Vui lòng giờ giây lát...',
            error_product_not_found: 'Không có sản phẩm nào tương ứng với tên miền "{domain}".',
        }
    }
}

Vue.use(VueI18n)

const i18n = new VueI18n({
    locale: 'en',
    messages: messages
})

export default i18n
