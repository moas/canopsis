/**
 * Mock for date now. Clear yourself after all tests.
 *
 * @param {number} nowTimestamp
 */
export const mockDateNow = (nowTimestamp) => {
  let dateNowSpy;

  beforeAll(() => {
    dateNowSpy = jest.spyOn(Date, 'now').mockReturnValue(nowTimestamp);
  });

  afterAll(() => {
    dateNowSpy.mockRestore();
  });
};

/**
 * Mock for requestAnimationFrame. Clear yourself after all tests.
 */
export const mockRequestAnimationFrame = () => {
  let requestAnimationFrameSpy = null;

  beforeEach(() => {
    requestAnimationFrameSpy = jest.spyOn(window, 'requestAnimationFrame')
      .mockImplementation(() => {});
  });

  afterEach(() => {
    requestAnimationFrameSpy.mockRestore();
  });
};

/**
 * Mock for date. Clear yourself after all tests.
 *
 * @param {number | Date} nowTimestamp
 */
export const mockDateGetTime = (nowTimestamp) => {
  let dateSpy;

  beforeAll(() => {
    dateSpy = jest
      .spyOn(Date.prototype, 'getTime')
      .mockReturnValue(nowTimestamp);
  });

  afterAll(() => {
    dateSpy.mockRestore();
  });
};

/**
 * Mock for the modals. Clear yourself after all tests.
 */
export const mockModals = () => {
  const modals = {
    show: jest.fn(),
    moduleName: 'modals',
  };

  afterEach(() => {
    modals.show.mockReset();
  });

  return modals;
};

/**
 * Mock for the popups. Clear yourself after all tests.
 */
export const mockPopups = () => {
  const popups = {
    error: jest.fn(),
  };

  afterEach(() => {
    popups.error.mockReset();
  });

  return popups;
};

/**
 * Mock for the sidebar. Clear yourself after all tests.
 */
export const mockSidebar = () => {
  const sidebar = {
    show: jest.fn(),
    hide: jest.fn(),
    moduleName: 'sidebar',
  };

  afterEach(() => {
    sidebar.show.mockReset();
    sidebar.hide.mockReset();
  });

  return sidebar;
};
