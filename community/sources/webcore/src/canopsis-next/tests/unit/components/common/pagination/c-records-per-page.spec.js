import { mount, createVueInstance } from '@unit/utils/vue';

import { PAGINATION_LIMIT, PAGINATION_PER_PAGE_VALUES } from '@/config';

import CRecordsPerPage from '@/components/common/pagination/c-records-per-page.vue';

const localVue = createVueInstance();

const stubs = {
  'v-select': {
    props: ['value', 'items'],
    template: `
      <select class="v-select" :value="value" @change="$listeners.input($event.target.value)">
        <option v-for="item in items" :key="item" :value="item">{{ item }}</option>
      </select>
    `,
  },
};

const factory = (options = {}) => mount(CRecordsPerPage, {
  localVue,
  stubs,
  ...options,
});

describe('c-records-per-page', () => {
  it('Default items is equal to PAGINATION_PER_PAGE_VALUES', () => {
    const wrapper = factory();

    expect(wrapper.vm.items).toEqual(PAGINATION_PER_PAGE_VALUES);
  });

  it('Default items property applied to select element', () => {
    const wrapper = factory();
    const select = wrapper.find('select.v-select');

    expect(select.vm.items).toEqual(PAGINATION_PER_PAGE_VALUES);
  });

  it('Custom items property applied to select element', () => {
    const items = [1, 2, 3];
    const wrapper = factory({ propsData: { items } });
    const select = wrapper.find('select.v-select');

    expect(select.vm.items).toEqual(items);
  });

  it('Default value property is equal to PAGINATION_LIMIT', () => {
    const wrapper = factory();

    expect(wrapper.vm.value).toBe(PAGINATION_LIMIT);
  });

  it('Default value property applied to select element', () => {
    const wrapper = factory();
    const select = wrapper.find('select.v-select');

    expect(select.vm.value).toBe(PAGINATION_LIMIT);
  });

  it('Custom value property applied to select element', () => {
    const value = PAGINATION_PER_PAGE_VALUES[0];
    const wrapper = factory({ propsData: { value } });
    const select = wrapper.find('select.v-select');

    expect(select.vm.value).toBe(value);
  });

  it('Custom items and value properties applied to select element', () => {
    const items = [1, 2, 3, 4, 5];
    const value = items[0];
    const wrapper = factory({ propsData: { value, items } });
    const select = wrapper.find('select.v-select');

    expect(select.vm.value).toBe(value);
    expect(select.vm.items).toEqual(items);
  });

  it('Set value into select element', () => {
    const value = PAGINATION_PER_PAGE_VALUES[0];
    const wrapper = factory();
    const select = wrapper.find('select.v-select');

    select.setValue(value);

    const inputEvents = wrapper.emitted('input');

    expect(inputEvents).toBeTruthy();
    expect(inputEvents).toHaveLength(1);
    expect(inputEvents[0].map(e => parseInt(e, 10))).toEqual([value]);
  });

  it('Renders `c-records-per-page` correctly', () => {
    const wrapper = mount(CRecordsPerPage, {
      localVue,
    });

    expect(wrapper.element).toMatchSnapshot();
  });
});
