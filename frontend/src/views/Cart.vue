<template>
  <div class="container">
    <h1 class="text-center">Корзина</h1>
    <CartTable />
    <div class="d-flex justify-content-end mt-3" v-if="$store.getters.numberOfItemsInCart">
      <button class="btn btn-primary" @click="buy" :disabled="loading">
        {{ loading ? 'Оформляем...' : 'Купить' }}
      </button>
    </div>
    
    <!-- Success Modal -->
    <SuccessModal 
      :show="showSuccessModal" 
      title="Заказ принят!"
      :message="successMessage"
      @close="showSuccessModal = false"
    />
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'
import { useStore } from 'vuex'
import CartTable from '@/components/cart/CartTable.vue';
import SuccessModal from '@/components/misc/SuccessModal.vue';

export default defineComponent({
  name: 'Cart',
  components: {
    CartTable,
    SuccessModal
  },
  setup() {
    const store = useStore();
    const loading = ref(false);
    const showSuccessModal = ref(false);
    const successMessage = ref('');
    
    async function buy() {
      if (loading.value) return;
      loading.value = true;
      try {
        const data = await store.dispatch('createOrder');
        successMessage.value = `Заказ номер: ` + (data?.id ?? '—') + ` принят. Хорошего дня! 😊`;
        showSuccessModal.value = true;
      } catch (e) {
        alert('Не удалось оформить заказ');
      } finally {
        loading.value = false;
      }
    }
    return { loading, buy, showSuccessModal, successMessage };
  }
})
</script>
