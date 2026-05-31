export default {
  common: {
    appName: 'GoRedisAdmin',
    language: 'Language',
    chinese: '中文',
    english: 'English',
    refresh: 'Refresh',
    save: 'Save',
    cancel: 'Cancel',
    close: 'Close',
    confirm: 'Confirm',
    warning: 'Notice',
    alert: 'Warning',
    operationSuccess: 'Operation completed'
  },
  header: {
    databases: 'Databases',
    redisInfo: 'Redis Info',
    logout: 'Sign out',
    github: 'GitHub'
  },
  login: {
    subtitle: 'Connect to the admin panel',
    username: 'Username',
    password: 'Password',
    usernamePlaceholder: 'Enter username',
    passwordPlaceholder: 'Enter password',
    loginButton: 'Sign In',
    validating: 'Validating...',
    success: 'Connected',
    redirecting: 'Entering the admin panel...',
    emptyCredentials: 'Username and password cannot be empty'
  },
  home: {
    databasesTitle: 'Databases',
    current: 'Current',
    selectDatabase: 'Please select a database',
    keyCount: 'keys',
    addKey: 'New Key',
    clearDatabase: 'Clear DB',
    flushAll: 'Flush All',
    flushAllConfirm: 'Are you sure you want to flush all databases?',
    batchDelete: 'Batch Delete',
    exportData: 'Export Data',
    searchPlaceholder: 'Search keys...',
    expireAt: 'TTL',
    type: 'Type',
    size: 'Size',
    actions: 'Actions',
    view: 'View',
    setExpire: 'Set Expire',
    delete: 'Delete',
    selectDatabaseHint: 'Select a database from the left to start',
    addKeyDialogTitle: 'New Key — {dbName}',
    viewDataTitle: 'View Data',
    selectDeleteWarning: 'Please select records to delete',
    batchDeleteConfirm: 'This operation will permanently delete selected records. Continue?',
    selectExportWarning: 'Please select keys to export',
    exportConfirm: 'Export selected data now?',
    exportSuccess: 'Export successful!',
    expirePromptMessage: 'Please enter expiration time (seconds, 0 means never expire)',
    expirePromptTitle: 'Set Expiration',
    deleteKeyConfirm: 'Delete [{key}]?',
    flushTargetAll: 'all databases',
    flushTargetDb: 'database {name}',
    flushConfirm: 'This operation will clear data in {target}. Continue?'
  },
  form: {
    keyName: 'Key',
    keyPlaceholder: 'Enter key name',
    type: 'Type',
    typePlaceholder: 'Select data type',
    expireTime: 'Expiration',
    expirePlaceholder: '0 = Never expire',
    score: 'Score',
    scorePlaceholder: 'Enter score',
    hashKey: 'Hash Key',
    hashKeyPlaceholder: 'Enter hash key name',
    value: 'Value',
    valuePlaceholder: 'Enter value',
    keyRequired: 'Please enter key',
    hashKeyRequired: 'Please enter hash key',
    typeRequired: 'Please select a type'
  },
  data: {
    value: 'Value',
    hashFields: 'Hash Fields',
    members: '{type} Members',
    index: '#',
    key: 'Key',
    unsupported: 'Preview is not available for this data type'
  },
  info: {
    title: 'Redis Server Info',
    waitLoading: 'Waiting for data...'
  },
  right: {
    home: 'Home',
    description: 'Page description'
  },
  request: {
    loginExpired: 'Session expired. Please sign in again.',
    responseNotArray: 'The response payload is not an array.',
    downloadFailed: 'File download failed:',
    error: 'Error'
  }
}
