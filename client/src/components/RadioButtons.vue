<template>
    <div :class="{ 'flex-col': orientation === 'vertical', 'flex-row': orientation === 'horizontal' }" class="flex gap-4">
        <label v-for="radio in radios" :for="radio.value" class="flex items-center gap-1.5 cursor-pointer text-sm">
            <input type="radio" name="radio-group" :id="radio.value" class="cursor-pointer text-black p-0 bg-white border-gray-300 outline-none ring-0 ring-transparent rounded-full form-input" :value="radio.value" @change="handleSelectRadio(radio.value)" />
            <span class="capitalize">{{radio.value}}</span>
        </label>
    </div>
</template>
<script lang="ts" setup>
import { PropType } from 'vue';

interface Radio {
    value: string;
}

const emit = defineEmits(['update:modelValue']);

defineProps({
    radios: {
        type: Array as PropType<Radio[]>,
        required: true
    },
    orientation: {
        type: String,
        required: false,
        validator: (value: string) => {
            return ['horizontal', 'vertical'].includes(value);
        },
        default: 'horizontal'
    }
})

const handleSelectRadio = (value: string) => {
    emit('update:modelValue', value);
}
</script>