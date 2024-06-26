let __databaseObj;
const __currentDatabaseVersion = 8;
const __constCachedItem = '__cached-items__';
const __databaseName = 'dessage-database'

function initDatabase() {
    return new Promise((resolve, reject) => {
        const request = indexedDB.open(__databaseName, __currentDatabaseVersion);

        request.onerror = function (event) {
            console.error("Database open failed:", event.target.error);
            reject(event.target.error);
        };

        request.onsuccess = function (event) {
            __databaseObj = event.target.result;
            console.log("Database open success, version=", __databaseObj.version);
            resolve(__databaseObj);
        };

        request.onupgradeneeded = function (event) {
            const db = event.target.result;
            if (!db.objectStoreNames.contains(__constCachedItem)) {
                const objectStore = db.createObjectStore(__constCachedItem, {keyPath: 'key'});
                objectStore.createIndex('keyIdx', 'key', {unique: true});
                console.log("Created cached item table successfully.");
            }
        };
    });
}

function closeDatabase() {
    if (__databaseObj) {
        __databaseObj.close();
        console.log("Database connection closed.");
    }
}

function databaseAddItem(storeName, data) {
    return new Promise((resolve, reject) => {
        const transaction = __databaseObj.transaction([storeName], 'readwrite');
        const objectStore = transaction.objectStore(storeName);
        const request = objectStore.add(data);

        request.onsuccess = () => {
            resolve(request.result);
        };

        request.onerror = event => {
            reject(`Error adding data to ${storeName}: ${event.target.error}`);
        };
    });
}

function databaseGetByIndex(storeName, idx, idxVal) {
    return new Promise((resolve, reject) => {
        try {
            // 启动事务
            const transaction = __databaseObj.transaction([storeName], 'readonly');
            // 获取对象存储
            const objectStore = transaction.objectStore(storeName);
            // 获取索引
            const index = objectStore.index(idx);  // 确保这是你创建的索引名

            // 使用索引查询数据
            const queryRequest = index.get(idxVal);  // key 是你想查找的键值

            queryRequest.onsuccess = function () {
                if (queryRequest.result) {
                    resolve(queryRequest.result);  // 返回找到的对象
                } else {
                    resolve(null);  // 没有找到对象
                }
            };

            queryRequest.onerror = function (event) {
                reject('Error in query by key: ' + event.target.error);
            };
        } catch (error) {
            reject('Transaction failed: ' + error.message);
        }
    });
}

function databaseGetByID(storeName, id) {
    return new Promise((resolve, reject) => {
        const transaction = __databaseObj.transaction([storeName], 'readonly');
        const objectStore = transaction.objectStore(storeName);

        const request = objectStore.get(id);

        request.onsuccess = event => {
            const result = event.target.result;
            if (result) {
                resolve(result);
            } else {
                resolve(null);
            }
        };

        request.onerror = event => {
            reject(`Error getting data from ${storeName}: ${event.target.error}`);
        };
    });
}

function databaseUpdate(storeName, id, newData) {
    return new Promise((resolve, reject) => {
        const transaction = __databaseObj.transaction([storeName], 'readwrite');
        const objectStore = transaction.objectStore(storeName);

        const request = objectStore.put({...newData, id});

        request.onsuccess = () => {
            resolve(`Data updated in ${storeName} successfully`);
        };

        request.onerror = event => {
            reject(`Error updating data in ${storeName}: ${event.target.error}`);
        };
    });
}

function databaseAddOrUpdate(storeName, data) {
    const transaction = __databaseObj.transaction([storeName], 'readwrite');
    const objectStore = transaction.objectStore(storeName);

    // Use put instead of add
    const request = objectStore.put(data);

    return new Promise((resolve, reject) => {
        request.onsuccess = () => {
            const isNewData = request.source === null;
            resolve({isNewData, id: request.result});
        };

        request.onerror = event => {
            reject(`Error adding/updating data in ${storeName}: ${event.target.error}`);
        };
    });
}

function databaseDelete(storeName, id) {
    return new Promise((resolve, reject) => {
        const transaction = __databaseObj.transaction([storeName], 'readwrite');
        const objectStore = transaction.objectStore(storeName);

        const request = objectStore.delete(id);

        request.onsuccess = () => {
            resolve(`Data deleted from ${storeName} successfully`);
        };

        request.onerror = event => {
            reject(`Error deleting data from ${storeName}: ${event.target.error}`);
        };
    });
}

function databaseDeleteByFilter(storeName, conditionFn) {
    return new Promise((resolve, reject) => {
        const transaction = __databaseObj.transaction([storeName], 'readwrite');
        const objectStore = transaction.objectStore(storeName);
        const request = objectStore.openCursor();

        request.onsuccess = event => {
            const cursor = event.target.result;
            if (cursor) {
                if (conditionFn(cursor.value)) {
                    cursor.delete();
                }
                cursor.continue();
            } else {
                resolve(`Data deleted from ${storeName} successfully`);
            }
        };

        request.onerror = event => {
            reject(`Error deleting data with condition from ${storeName}: ${event.target.error}`);
        };
    });
}

function databaseQueryAll(storeName) {
    return new Promise((resolve, reject) => {
        const transaction = __databaseObj.transaction([storeName], 'readonly');
        const objectStore = transaction.objectStore(storeName);
        const request = objectStore.getAll();

        request.onsuccess = event => {
            const data = event.target.result;
            resolve(data);
        };

        request.onerror = event => {
            reject(`Error getting all data from ${storeName}: ${event.target.error}`);
        };
    });
}

function databaseQueryByFilter(storeName, conditionFn) {
    return new Promise((resolve, reject) => {
        const transaction = __databaseObj.transaction([storeName], 'readonly');
        const objectStore = transaction.objectStore(storeName);
        const request = objectStore.openCursor();

        const results = [];

        request.onsuccess = event => {
            const cursor = event.target.result;
            if (cursor) {
                const data = cursor.value;
                if (conditionFn(data)) {
                    results.push(data);
                }
                cursor.continue();
            } else {
                resolve(results);
            }
        };

        request.onerror = event => {
            reject(`Error querying data from ${storeName}: ${event.target.error}`);
        };
    });
}

function databaseCleanByFilter(storeName, newData, conditionFn) {
    const clearAndFillTransaction = __databaseObj.transaction([storeName], 'readwrite');
    const objectStore = clearAndFillTransaction.objectStore(storeName);

    // 根据条件删除数据
    const clearRequest = objectStore.openCursor();

    clearRequest.onsuccess = event => {
        const cursor = event.target.result;
        if (cursor) {
            const data = cursor.value;
            if (conditionFn(data)) {
                cursor.delete();
            }
            cursor.continue();
        } else {
            const fillTransaction = __databaseObj.transaction([storeName], 'readwrite');
            const fillObjectStore = fillTransaction.objectStore(storeName);

            if (Array.isArray(newData) && newData.length > 0) {
                newData.forEach(data => {
                    if (!data.id) {
                        fillObjectStore.add(data);
                    } else {
                        fillObjectStore.put(data);
                    }
                });
            }

            fillTransaction.oncomplete = () => {
                console.log(`Table ${storeName} cleared and filled with new data.`);
            };

            fillTransaction.onerror = event => {
                console.error(`Error filling table ${storeName}: ${event.target.error}`);
            };
        }
    };

    clearRequest.onerror = event => {
        console.error(`Error clearing table ${storeName}: ${event.target.error}`);
    };
}

function databaseDeleteTable(tableName) {

    const request = indexedDB.open(__databaseName);

    request.onsuccess = function (event) {
        const db = event.target.result;
        const transaction = db.transaction(tableName, 'readwrite');
        const objectStore = transaction.objectStore(tableName);

        const clearRequest = objectStore.clear();

        clearRequest.onsuccess = function () {
            console.log(`${tableName} has been cleared`);
        };

        clearRequest.onerror = function (event) {
            console.error('Clear object store error:', event.target.error);
        };
    };

    request.onerror = function (event) {
        console.error('Database error:', event.target.error);
    };
}