// +build windows

package win

const (
	// Define the severity codes

	// The operation completed successfully.
	ERROR_SUCCESS = 0

	NO_ERROR = 0 // dderror

	// Incorrect function.
	ERROR_INVALID_FUNCTION = 1 // dderror

	// The system cannot find the file specified.
	ERROR_FILE_NOT_FOUND = 2

	// The system cannot find the path specified.
	ERROR_PATH_NOT_FOUND = 3

	// The system cannot open the file.
	ERROR_TOO_MANY_OPEN_FILES = 4

	// Access is denied.
	ERROR_ACCESS_DENIED = 5

	// The handle is invalid.
	ERROR_INVALID_HANDLE = 6

	// The storage control blocks were destroyed.
	ERROR_ARENA_TRASHED = 7

	// Not enough storage is available to process this command.
	ERROR_NOT_ENOUGH_MEMORY = 8 // dderror

	// The storage control block address is invalid.
	ERROR_INVALID_BLOCK = 9

	// The environment is incorrect.
	ERROR_BAD_ENVIRONMENT = 10

	// An attempt was made to load a program with an incorrect format.
	ERROR_BAD_FORMAT = 11

	// The access code is invalid.
	ERROR_INVALID_ACCESS = 12

	// The data is invalid.
	ERROR_INVALID_DATA = 13

	// Not enough storage is available to complete this operation.
	ERROR_OUTOFMEMORY = 14

	// The system cannot find the drive specified.
	ERROR_INVALID_DRIVE = 15

	// The directory cannot be removed.
	ERROR_CURRENT_DIRECTORY = 0x10

	// The system cannot move the file
	// to a different disk drive.
	ERROR_NOT_SAME_DEVICE = 17

	// There are no more files.
	ERROR_NO_MORE_FILES = 18

	// The media is write protected.
	ERROR_WRITE_PROTECT = 19

	// The system cannot find the device specified.
	ERROR_BAD_UNIT = 20

	// The device is not ready.
	ERROR_NOT_READY = 21

	// The device does not recognize the command.
	ERROR_BAD_COMMAND = 22

	// Data error (cyclic redundancy check)
	ERROR_CRC = 23

	// The program issued a command but the command length is incorrect.
	ERROR_BAD_LENGTH = 24

	// The drive cannot locate a specific area or track on the disk.
	ERROR_SEEK = 25

	// The specified disk or diskette cannot be accessed.
	ERROR_NOT_DOS_DISK = 26

	// The drive cannot find the sector requested.
	ERROR_SECTOR_NOT_FOUND = 27

	// The printer is out of paper.
	ERROR_OUT_OF_PAPER = 28

	// The system cannot write to the specified device.
	ERROR_WRITE_FAULT = 29

	// The system cannot read from the specified device.
	ERROR_READ_FAULT = 30

	// A device attached to the system is not functioning.
	ERROR_GEN_FAILURE = 31

	// The process cannot access the file because it is being used by another process.
	ERROR_SHARING_VIOLATION = 0x20

	// The process cannot access the file because
	// another process has locked a portion of the file.
	ERROR_LOCK_VIOLATION = 33

	// The wrong diskette is in the drive. Insert %2 (Volume Serial Number: %3)
	// into drive %1.
	ERROR_WRONG_DISK = 34

	// Too many files opened for sharing.
	ERROR_SHARING_BUFFER_EXCEEDED = 36

	// Reached end of file.
	ERROR_HANDLE_EOF = 38

	// The disk is full.
	ERROR_HANDLE_DISK_FULL = 39

	// The network request is not supported.
	ERROR_NOT_SUPPORTED = 50

	// The remote computer is not available.
	ERROR_REM_NOT_LIST = 51

	// A duplicate name exists on the network.
	ERROR_DUP_NAME = 52

	// The network path was not found.
	ERROR_BAD_NETPATH = 53

	// The network is busy.
	ERROR_NETWORK_BUSY = 54

	// The specified network resource or device is no longer
	// available.
	ERROR_DEV_NOT_EXIST = 55 // dderror

	// The network BIOS command limit has been reached.
	ERROR_TOO_MANY_CMDS = 56

	// A network adapter hardware error occurred.
	ERROR_ADAP_HDW_ERR = 57

	// The specified server cannot perform the requested
	// operation.
	ERROR_BAD_NET_RESP = 58

	// An unexpected network error occurred.
	ERROR_UNEXP_NET_ERR = 59

	// The remote adapter is not compatible.
	ERROR_BAD_REM_ADAP = 60

	// The printer queue is full.
	ERROR_PRINTQ_FULL = 61

	// Space to store the file waiting to be printed is
	// not available on the server.
	ERROR_NO_SPOOL_SPACE = 62

	// Your file waiting to be printed was deleted.
	ERROR_PRINT_CANCELLED = 63

	// The specified network name is no longer available.
	ERROR_NETNAME_DELETED = 0x40

	// Network access is denied.
	ERROR_NETWORK_ACCESS_DENIED = 65

	// The network resource type is not correct.
	ERROR_BAD_DEV_TYPE = 66

	// The network name cannot be found.
	ERROR_BAD_NET_NAME = 67

	// The name limit for the local computer network
	// adapter card was exceeded.
	ERROR_TOO_MANY_NAMES = 68

	// The network BIOS session limit was exceeded.
	ERROR_TOO_MANY_SESS = 69

	// The remote server has been paused or is in the
	// process of being started.
	ERROR_SHARING_PAUSED = 70

	// No more connections can be made to this remote computer at this time
	// because there are already as many connections as the computer can accept.
	ERROR_REQ_NOT_ACCEP = 71

	// The specified printer or disk device has been paused.
	ERROR_REDIR_PAUSED = 72

	// The file exists.
	ERROR_FILE_EXISTS = 80

	// The directory or file cannot be created.
	ERROR_CANNOT_MAKE = 82

	// Fail on INT 24
	ERROR_FAIL_I24 = 83

	// Storage to process this request is not available.
	ERROR_OUT_OF_STRUCTURES = 84

	// The local device name is already in use.
	ERROR_ALREADY_ASSIGNED = 85

	// The specified network password is not correct.
	ERROR_INVALID_PASSWORD = 86

	// The parameter is incorrect.
	ERROR_INVALID_PARAMETER = 87 // dderror

	// A write fault occurred on the network.
	ERROR_NET_WRITE_FAULT = 88

	// The system cannot start another process at
	// this time.
	ERROR_NO_PROC_SLOTS = 89

	// Cannot create another system semaphore.
	ERROR_TOO_MANY_SEMAPHORES = 100

	// The exclusive semaphore is owned by another process.
	ERROR_EXCL_SEM_ALREADY_OWNED = 101

	// The semaphore is set and cannot be closed.
	ERROR_SEM_IS_SET = 102

	// The semaphore cannot be set again.
	ERROR_TOO_MANY_SEM_REQUESTS = 103

	// Cannot request exclusive semaphores at interrupt time.
	ERROR_INVALID_AT_INTERRUPT_TIME = 104

	// The previous ownership of this semaphore has ended.
	ERROR_SEM_OWNER_DIED = 105

	// Insert the diskette for drive %1.
	ERROR_SEM_USER_LIMIT = 106

	// Program stopped because alternate diskette was not inserted.
	ERROR_DISK_CHANGE = 107

	// The disk is in use or locked by
	// another process.
	ERROR_DRIVE_LOCKED = 108

	// The pipe has been ended.
	ERROR_BROKEN_PIPE = 109

	// The system cannot open the device or file specified.
	ERROR_OPEN_FAILED = 110

	// The file name is too long.
	ERROR_BUFFER_OVERFLOW = 111

	// There is not enough space on the disk.
	ERROR_DISK_FULL = 112

	// No more internal file identifiers available.
	ERROR_NO_MORE_SEARCH_HANDLES = 113

	// The target internal file identifier is incorrect.
	ERROR_INVALID_TARGET_HANDLE = 114

	// The IOCTL call made by the application program is not correct.
	ERROR_INVALID_CATEGORY = 117

	// The verify-on-write switch parameter value is not correct.
	ERROR_INVALID_VERIFY_SWITCH = 118

	// The system does not support the command requested.
	ERROR_BAD_DRIVER_LEVEL = 119

	// This function is only valid in Windows NT mode.
	ERROR_CALL_NOT_IMPLEMENTED = 120

	// The semaphore timeout period has expired.
	ERROR_SEM_TIMEOUT = 121

	// The data area passed to a system call is too small.
	ERROR_INSUFFICIENT_BUFFER = 122 // dderror

	// The filename, directory name, or volume label syntax is incorrect.
	ERROR_INVALID_NAME = 123

	// The system call level is not correct.
	ERROR_INVALID_LEVEL = 124

	// The disk has no volume label.
	ERROR_NO_VOLUME_LABEL = 125

	// The specified module could not be found.
	ERROR_MOD_NOT_FOUND = 126

	// The specified procedure could not be found.
	ERROR_PROC_NOT_FOUND = 127

	// There are no child processes to wait for.
	ERROR_WAIT_NO_CHILDREN = 0x80

	// The %1 application cannot be run in Windows NT mode.
	ERROR_CHILD_NOT_COMPLETE = 129

	// Attempt to use a file handle to an open disk partition for an
	// operation other than raw disk I/O.
	ERROR_DIRECT_ACCESS_HANDLE = 130

	// An attempt was made to move the file pointer before the beginning of the file.
	ERROR_NEGATIVE_SEEK = 131

	// The file pointer cannot be set on the specified device or file.
	ERROR_SEEK_ON_DEVICE = 132

	// A JOIN or SUBST command
	// cannot be used for a drive that
	// contains previously joined drives.
	ERROR_IS_JOIN_TARGET = 133

	// An attempt was made to use a
	// JOIN or SUBST command on a drive that has
	// already been joined.
	ERROR_IS_JOINED = 134

	// An attempt was made to use a
	// JOIN or SUBST command on a drive that has
	// already been substituted.
	ERROR_IS_SUBSTED = 135

	// The system tried to delete
	// the JOIN of a drive that is not joined.
	ERROR_NOT_JOINED = 136

	// The system tried to delete the
	// substitution of a drive that is not substituted.
	ERROR_NOT_SUBSTED = 137

	// The system tried to join a drive to a directory on a joined drive.
	ERROR_JOIN_TO_JOIN = 138

	// The system tried to substitute a drive to a directory on a substituted drive.
	ERROR_SUBST_TO_SUBST = 139

	// The system tried to join a drive to a directory on a substituted drive.
	ERROR_JOIN_TO_SUBST = 140

	// The system tried to SUBST a drive to a directory on a joined drive.
	ERROR_SUBST_TO_JOIN = 141

	// The system cannot perform a JOIN or SUBST at this time.
	ERROR_BUSY_DRIVE = 142

	// The system cannot join or substitute a
	// drive to or for a directory on the same drive.
	ERROR_SAME_DRIVE = 143

	// The directory is not a subdirectory of the root directory.
	ERROR_DIR_NOT_ROOT = 144

	// The directory is not empty.
	ERROR_DIR_NOT_EMPTY = 145

	// The path specified is being used in a substitute.
	ERROR_IS_SUBST_PATH = 146

	// Not enough resources are available to process this command.
	ERROR_IS_JOIN_PATH = 147

	// The path specified cannot be used at this time.
	ERROR_PATH_BUSY = 148

	// An attempt was made to join or substitute a drive for which a directory
	// on the drive is the target of a previous substitute.
	ERROR_IS_SUBST_TARGET = 149

	// System trace information was not specified in your
	// CONFIG.SYS file, or tracing is disallowed.
	ERROR_SYSTEM_TRACE = 150

	// The number of specified semaphore events for
	// DosMuxSemWait is not correct.
	ERROR_INVALID_EVENT_COUNT = 151

	// DosMuxSemWait did not execute; too many semaphores
	// are already set.
	ERROR_TOO_MANY_MUXWAITERS = 152

	// The DosMuxSemWait list is not correct.
	ERROR_INVALID_LIST_FORMAT = 153

	//  The volume label you entered exceeds the label character
	//  limit of the target file system.
	ERROR_LABEL_TOO_LONG = 154

	// Cannot create another thread.
	ERROR_TOO_MANY_TCBS = 155

	// The recipient process has refused the signal.
	ERROR_SIGNAL_REFUSED = 156

	// The segment is already discarded and cannot be locked.
	ERROR_DISCARDED = 157

	// The segment is already unlocked.
	ERROR_NOT_LOCKED = 158

	// The address for the thread ID is not correct.
	ERROR_BAD_THREADID_ADDR = 159

	// The argument string passed to DosExecPgm is not correct.
	ERROR_BAD_ARGUMENTS = 160

	// The specified path is invalid.
	ERROR_BAD_PATHNAME = 161

	// A signal is already pending.
	ERROR_SIGNAL_PENDING = 162

	// No more threads can be created in the system.
	ERROR_MAX_THRDS_REACHED = 164

	// Unable to lock a region of a file.
	ERROR_LOCK_FAILED = 167

	// The requested resource is in use.
	ERROR_BUSY = 170

	// A lock request was not outstanding for the supplied cancel region.
	ERROR_CANCEL_VIOLATION = 173

	// The file system does not support atomic changes to the lock type.
	ERROR_ATOMIC_LOCKS_NOT_SUPPORTED = 174

	// The system detected a segment number that was not correct.
	ERROR_INVALID_SEGMENT_NUMBER = 180

	// The operating system cannot run %1.
	ERROR_INVALID_ORDINAL = 182

	// Cannot create a file when that file already exists.
	ERROR_ALREADY_EXISTS = 183

	// The flag passed is not correct.
	ERROR_INVALID_FLAG_NUMBER = 186

	// The specified system semaphore name was not found.
	ERROR_SEM_NOT_FOUND = 187

	// The operating system cannot run %1.
	ERROR_INVALID_STARTING_CODESEG = 188

	// The operating system cannot run %1.
	ERROR_INVALID_STACKSEG = 189

	// The operating system cannot run %1.
	ERROR_INVALID_MODULETYPE = 190

	// Cannot run %1 in Windows NT mode.
	ERROR_INVALID_EXE_SIGNATURE = 191

	// The operating system cannot run %1.
	ERROR_EXE_MARKED_INVALID = 192

	// %1 is not a valid Windows NT application.
	ERROR_BAD_EXE_FORMAT = 193

	// The operating system cannot run %1.
	ERROR_ITERATED_DATA_EXCEEDS_64k = 194

	// The operating system cannot run %1.
	ERROR_INVALID_MINALLOCSIZE = 195

	// The operating system cannot run this application program.
	ERROR_DYNLINK_FROM_INVALID_RING = 196

	// The operating system is not presently configured to run this application.
	ERROR_IOPL_NOT_ENABLED = 197

	// The operating system cannot run %1.
	ERROR_INVALID_SEGDPL = 198

	// The operating system cannot run this
	// application program.
	ERROR_AUTODATASEG_EXCEEDS_64k = 199

	// The code segment cannot be greater than or equal to 64KB.
	ERROR_RING2SEG_MUST_BE_MOVABLE = 200

	// The operating system cannot run %1.
	ERROR_RELOC_CHAIN_XEEDS_SEGLIM = 201

	// The operating system cannot run %1.
	ERROR_INFLOOP_IN_RELOC_CHAIN = 202

	// The system could not find the environment option that was entered.
	ERROR_ENVVAR_NOT_FOUND = 203

	// No process in the command subtree has a signal handler.
	ERROR_NO_SIGNAL_SENT = 205

	// The filename or extension is too long.
	ERROR_FILENAME_EXCED_RANGE = 206

	// The ring 2 stack is in use.
	ERROR_RING2_STACK_IN_USE = 207

	// The global filename characters, * or ?, are entered
	// incorrectly or too many global filename characters are specified.
	ERROR_META_EXPANSION_TOO_LONG = 208

	// The signal being posted is not correct.
	ERROR_INVALID_SIGNAL_NUMBER = 209

	// The signal handler cannot be set.
	ERROR_THREAD_1_INACTIVE = 210

	// The segment is locked and cannot be reallocated.
	ERROR_LOCKED = 212

	// Too many dynamic link modules are attached to this
	// program or dynamic link module.
	ERROR_TOO_MANY_MODULES = 214

	// Can't nest calls to LoadModule.
	ERROR_NESTING_NOT_ALLOWED = 215

	//  The image file %1 is valid, but is for a machine type other
	//  than the current machine.
	ERROR_EXE_MACHINE_TYPE_MISMATCH = 216

	// The pipe state is invalid.
	ERROR_BAD_PIPE = 230

	// All pipe instances are busy.
	ERROR_PIPE_BUSY = 231

	// The pipe is being closed.
	ERROR_NO_DATA = 232

	// No process is on the other end of the pipe.
	ERROR_PIPE_NOT_CONNECTED = 233

	// More data is available.
	ERROR_MORE_DATA = 234 // dderror

	// The session was cancelled.
	ERROR_VC_DISCONNECTED = 240

	// The specified extended attribute name was invalid.
	ERROR_INVALID_EA_NAME = 254

	// The extended attributes are inconsistent.
	ERROR_EA_LIST_INCONSISTENT = 255

	// No more data is available.
	ERROR_NO_MORE_ITEMS = 259

	// The Copy API cannot be used.
	ERROR_CANNOT_COPY = 266

	// The directory name is invalid.
	ERROR_DIRECTORY = 267

	// The extended attributes did not fit in the buffer.
	ERROR_EAS_DIDNT_FIT = 275

	// The extended attribute file on the mounted file system is corrupt.
	ERROR_EA_FILE_CORRUPT = 276

	// The extended attribute table file is full.
	ERROR_EA_TABLE_FULL = 277

	// The specified extended attribute handle is invalid.
	ERROR_INVALID_EA_HANDLE = 278

	// The mounted file system does not support extended attributes.
	ERROR_EAS_NOT_SUPPORTED = 282

	// Attempt to release mutex not owned by caller.
	ERROR_NOT_OWNER = 288

	// Too many posts were made to a semaphore.
	ERROR_TOO_MANY_POSTS = 298

	// Only part of a Read/WriteProcessMemory request was completed.
	ERROR_PARTIAL_COPY = 299

	// The system cannot find message for message number $%1
	// in message file for %2.
	ERROR_MR_MID_NOT_FOUND = 317

	// Attempt to access invalid address.
	ERROR_INVALID_ADDRESS = 487

	// Arithmetic result exceeded 32 bits.
	ERROR_ARITHMETIC_OVERFLOW = 534

	// There is a process on other end of the pipe.
	ERROR_PIPE_CONNECTED = 535

	// Waiting for a process to open the other end of the pipe.
	ERROR_PIPE_LISTENING = 536

	// Access to the extended attribute was denied.
	ERROR_EA_ACCESS_DENIED = 994

	// The I/O operation has been aborted because of either a thread exit
	// or an application request.
	ERROR_OPERATION_ABORTED = 995

	// Overlapped I/O event is not in a signalled state.
	ERROR_IO_INCOMPLETE = 996

	// Overlapped I/O operation is in progress.
	ERROR_IO_PENDING = 997 // dderror

	// Invalid access to memory location.
	ERROR_NOACCESS = 998

	// Error performing inpage operation.
	ERROR_SWAPERROR = 999

	// Recursion too deep, stack overflowed.
	ERROR_STACK_OVERFLOW = 1001

	// The window cannot act on the sent message.
	ERROR_INVALID_MESSAGE = 1002

	// Cannot complete this function.
	ERROR_CAN_NOT_COMPLETE = 1003

	// Invalid flags.
	ERROR_INVALID_FLAGS = 1004

	// The volume does not contain a recognized file system.
	// Please make sure that all required file system drivers are loaded and that the
	// volume is not corrupt.
	ERROR_UNRECOGNIZED_VOLUME = 1005

	// The volume for a file has been externally altered such that the
	// opened file is no longer valid.
	ERROR_FILE_INVALID = 1006

	// The requested operation cannot be performed in full-screen mode.
	ERROR_FULLSCREEN_MODE = 1007

	// An attempt was made to reference a token that does not exist.
	ERROR_NO_TOKEN = 1008

	// The configuration registry database is corrupt.
	ERROR_BADDB = 1009

	// The configuration registry key is invalid.
	ERROR_BADKEY = 1010

	// The configuration registry key could not be opened.
	ERROR_CANTOPEN = 1011

	// The configuration registry key could not be read.
	ERROR_CANTREAD = 1012

	// The configuration registry key could not be written.
	ERROR_CANTWRITE = 1013

	// One of the files in the Registry database had to be recovered
	// by use of a log or alternate copy.  The recovery was successful.
	ERROR_REGISTRY_RECOVERED = 1014

	// The Registry is corrupt. The structure of one of the files that contains
	// Registry data is corrupt, or the system's image of the file in memory
	// is corrupt, or the file could not be recovered because the alternate
	// copy or log was absent or corrupt.
	ERROR_REGISTRY_CORRUPT = 1015

	// An I/O operation initiated by the Registry failed unrecoverably.
	// The Registry could not read in, or write out, or flush, one of the files
	// that contain the system's image of the Registry.
	ERROR_REGISTRY_IO_FAILED = 1016

	// The system has attempted to load or restore a file into the Registry, but the
	// specified file is not in a Registry file format.
	ERROR_NOT_REGISTRY_FILE = 1017

	// Illegal operation attempted on a Registry key which has been marked for deletion.
	ERROR_KEY_DELETED = 1018

	// System could not allocate the required space in a Registry log.
	ERROR_NO_LOG_SPACE = 1019

	// Cannot create a symbolic link in a Registry key that already
	// has subkeys or values.
	ERROR_KEY_HAS_CHILDREN = 1020

	// Cannot create a stable subkey under a volatile parent key.
	ERROR_CHILD_MUST_BE_VOLATILE = 1021

	// A notify change request is being completed and the information
	// is not being returned in the caller's buffer. The caller now
	// needs to enumerate the files to find the changes.
	ERROR_NOTIFY_ENUM_DIR = 1022

	// A stop control has been sent to a service which other running services
	// are dependent on.
	ERROR_DEPENDENT_SERVICES_RUNNING = 1051

	// The requested control is not valid for this service
	ERROR_INVALID_SERVICE_CONTROL = 1052

	// The service did not respond to the start or control request in a timely
	// fashion.
	ERROR_SERVICE_REQUEST_TIMEOUT = 1053

	// A thread could not be created for the service.
	ERROR_SERVICE_NO_THREAD = 1054

	// The service database is locked.
	ERROR_SERVICE_DATABASE_LOCKED = 1055

	// An instance of the service is already running.
	ERROR_SERVICE_ALREADY_RUNNING = 1056

	// The account name is invalid or does not exist.
	ERROR_INVALID_SERVICE_ACCOUNT = 1057

	// The specified service is disabled and cannot be started.
	ERROR_SERVICE_DISABLED = 1058

	// Circular service dependency was specified.
	ERROR_CIRCULAR_DEPENDENCY = 1059

	// The specified service does not exist as an installed service.
	ERROR_SERVICE_DOES_NOT_EXIST = 1060

	// The service cannot accept control messages at this time.
	ERROR_SERVICE_CANNOT_ACCEPT_CTRL = 1061

	// The service has not been started.
	ERROR_SERVICE_NOT_ACTIVE = 1062

	// The service process could not connect to the service controller.
	ERROR_FAILED_SERVICE_CONTROLLER_ = 1063

	// An exception occurred in the service when handling the control request.
	ERROR_EXCEPTION_IN_SERVICE = 1064

	// The database specified does not exist.
	ERROR_DATABASE_DOES_NOT_EXIST = 1065

	// The service has returned a service-specific error code.
	ERROR_SERVICE_SPECIFIC_ERROR = 1066

	// The process terminated unexpectedly.
	ERROR_PROCESS_ABORTED = 1067

	// The dependency service or group failed to start.
	ERROR_SERVICE_DEPENDENCY_FAIL = 1068

	// The service did not start due to a logon failure.
	ERROR_SERVICE_LOGON_FAILED = 1069

	// After starting, the service hung in a start-pending state.
	ERROR_SERVICE_START_HANG = 1070

	// The specified service database lock is invalid.
	ERROR_INVALID_SERVICE_LOCK = 1071

	// The specified service has been marked for deletion.
	ERROR_SERVICE_MARKED_FOR_DELETE = 1072

	// The specified service already exists.
	ERROR_SERVICE_EXISTS = 1073

	// The system is currently running with the last-known-good configuration.
	ERROR_ALREADY_RUNNING_LKG = 1074

	// The dependency service does not exist or has been marked for
	// deletion.
	ERROR_SERVICE_DEPENDENCY_DELETED = 1075

	// The current boot has already been accepted for use as the
	// last-known-good control set.
	ERROR_BOOT_ALREADY_ACCEPTED = 1076

	// No attempts to start the service have been made since the last boot.
	ERROR_SERVICE_NEVER_STARTED = 1077

	// The name is already in use as either a service name or a service display
	// name.
	ERROR_DUPLICATE_SERVICE_NAME = 1078

	//  The account specified for this service is different from the account
	//  specified for other services running in the same process.
	ERROR_DIFFERENT_SERVICE_ACCOUNT = 1079

	// The physical end of the tape has been reached.
	ERROR_END_OF_MEDIA = 1100

	// A tape access reached a filemark.
	ERROR_FILEMARK_DETECTED = 1101

	// Beginning of tape or partition was encountered.
	ERROR_BEGINNING_OF_MEDIA = 1102

	// A tape access reached the end of a set of files.
	ERROR_SETMARK_DETECTED = 1103

	// No more data is on the tape.
	ERROR_NO_DATA_DETECTED = 1104

	// Tape could not be partitioned.
	ERROR_PARTITION_FAILURE = 1105

	// When accessing a new tape of a multivolume partition, the current
	// blocksize is incorrect.
	ERROR_INVALID_BLOCK_LENGTH = 1106

	// Tape partition information could not be found when loading a tape.
	ERROR_DEVICE_NOT_PARTITIONED = 1107

	// Unable to lock the media eject mechanism.
	ERROR_UNABLE_TO_LOCK_MEDIA = 1108

	// Unable to unload the media.
	ERROR_UNABLE_TO_UNLOAD_MEDIA = 1109

	// Media in drive may have changed.
	ERROR_MEDIA_CHANGED = 1110

	// The I/O bus was reset.
	ERROR_BUS_RESET = 1111

	// No media in drive.
	ERROR_NO_MEDIA_IN_DRIVE = 1112

	// No mapping for the Unicode character exists in the target multi-byte code page.
	ERROR_NO_UNICODE_TRANSLATION = 1113

	// A dynamic link library (DLL) initialization routine failed.
	ERROR_DLL_INIT_FAILED = 1114

	// A system shutdown is in progress.
	ERROR_SHUTDOWN_IN_PROGRESS = 1115

	// Unable to abort the system shutdown because no shutdown was in progress.
	ERROR_NO_SHUTDOWN_IN_PROGRESS = 1116

	// The request could not be performed because of an I/O device error.
	ERROR_IO_DEVICE = 1117

	// No serial device was successfully initialized.  The serial driver will unload.
	ERROR_SERIAL_NO_DEVICE = 1118

	// Unable to open a device that was sharing an interrupt request (IRQ)
	// with other devices. At least one other device that uses that IRQ
	// was already opened.
	ERROR_IRQ_BUSY = 1119

	// A serial I/O operation was completed by another write to the serial port.
	// (The IOCTL_SERIAL_XOFF_COUNTER reached zero.)
	ERROR_MORE_WRITES = 1120

	// A serial I/O operation completed because the time-out period expired.
	// (The IOCTL_SERIAL_XOFF_COUNTER did not reach zero.)
	ERROR_COUNTER_TIMEOUT = 1121

	// No ID address mark was found on the floppy disk.
	ERROR_FLOPPY_ID_MARK_NOT_FOUND = 1122

	// Mismatch between the floppy disk sector ID field and the floppy disk
	// controller track address.
	ERROR_FLOPPY_WRONG_CYLINDER = 1123

	// The floppy disk controller reported an error that is not recognized
	// by the floppy disk driver.
	ERROR_FLOPPY_UNKNOWN_ERROR = 1124

	// The floppy disk controller returned inconsistent results in its registers.
	ERROR_FLOPPY_BAD_REGISTERS = 1125

	// While accessing the hard disk, a recalibrate operation failed, even after retries.
	ERROR_DISK_RECALIBRATE_FAILED = 1126

	// While accessing the hard disk, a disk operation failed even after retries.
	ERROR_DISK_OPERATION_FAILED = 1127

	// While accessing the hard disk, a disk controller reset was needed, but
	// even that failed.
	ERROR_DISK_RESET_FAILED = 1128

	// Physical end of tape encountered.
	ERROR_EOM_OVERFLOW = 1129

	// Not enough server storage is available to process this command.
	ERROR_NOT_ENOUGH_SERVER_MEMORY = 1130

	// A potential deadlock condition has been detected.
	ERROR_POSSIBLE_DEADLOCK = 1131

	// The base address or the file offset specified does not have the proper
	// alignment.
	ERROR_MAPPED_ALIGNMENT = 1132

	// An attempt to change the system power state was vetoed by another
	// application or driver.
	ERROR_SET_POWER_STATE_VETOED = 1140

	// The system BIOS failed an attempt to change the system power state.
	ERROR_SET_POWER_STATE_FAILED = 1141

	//  An attempt was made to create more links on a file than
	//  the file system supports.
	ERROR_TOO_MANY_LINKS = 1142

	// The specified program requires a newer version of Windows.
	ERROR_OLD_WIN_VERSION = 1150

	// The specified program is not a Windows or MS-DOS program.
	ERROR_APP_WRONG_OS = 1151

	// Cannot start more than one instance of the specified program.
	ERROR_SINGLE_INSTANCE_APP = 1152

	//  The specified program was written for an older version of Windows.
	ERROR_RMODE_APP = 1153

	// One of the library files needed to run this application is damaged.
	ERROR_INVALID_DLL = 1154

	// No application is associated with the specified file for this operation.
	ERROR_NO_ASSOCIATION = 1155

	// An error occurred in sending the command to the application.
	ERROR_DDE_FAIL = 1156

	// One of the library files needed to run this application cannot be found.
	ERROR_DLL_NOT_FOUND = 1157

	// Winnet32 Status Codes
	// The specified username is invalid.
	ERROR_BAD_USERNAME = 2202

	// This network connection does not exist.
	ERROR_NOT_CONNECTED = 2250

	// This network connection has files open or requests pending.
	ERROR_OPEN_FILES = 2401

	// Active connections still exist.
	ERROR_ACTIVE_CONNECTIONS = 2402

	// The device is in use by an active process and cannot be disconnected.
	ERROR_DEVICE_IN_USE = 2404

	// The specified device name is invalid.
	ERROR_BAD_DEVICE = 1200

	// The device is not currently connected but it is a remembered connection.
	ERROR_CONNECTION_UNAVAIL = 1201

	// An attempt was made to remember a device that had previously been remembered.
	ERROR_DEVICE_ALREADY_REMEMBERED = 1202

	// No network provider accepted the given network path.
	ERROR_NO_NET_OR_BAD_PATH = 1203

	// The specified network provider name is invalid.
	ERROR_BAD_PROVIDER = 1204

	// Unable to open the network connection profile.
	ERROR_CANNOT_OPEN_PROFILE = 1205

	// The network connection profile is corrupt.
	ERROR_BAD_PROFILE = 1206

	// Cannot enumerate a non-container.
	ERROR_NOT_CONTAINER = 1207

	// An extended error has occurred.
	ERROR_EXTENDED_ERROR = 1208

	// The format of the specified group name is invalid.
	ERROR_INVALID_GROUPNAME = 1209

	// The format of the specified computer name is invalid.
	ERROR_INVALID_COMPUTERNAME = 1210

	// The format of the specified event name is invalid.
	ERROR_INVALID_EVENTNAME = 1211

	// The format of the specified domain name is invalid.
	ERROR_INVALID_DOMAINNAME = 1212

	// The format of the specified service name is invalid.
	ERROR_INVALID_SERVICENAME = 1213

	// The format of the specified network name is invalid.
	ERROR_INVALID_NETNAME = 1214

	// The format of the specified share name is invalid.
	ERROR_INVALID_SHARENAME = 1215

	// The format of the specified password is invalid.
	ERROR_INVALID_PASSWORDNAME = 1216

	// The format of the specified message name is invalid.
	ERROR_INVALID_MESSAGENAME = 1217

	// The format of the specified message destination is invalid.
	ERROR_INVALID_MESSAGEDEST = 1218

	// The credentials supplied conflict with an existing set of credentials.
	ERROR_SESSION_CREDENTIAL_CONFLICT = 1219

	// An attempt was made to establish a session to a network server, but there
	// are already too many sessions established to that server.
	ERROR_REMOTE_SESSION_LIMIT_EXCEEDED = 1220

	// The workgroup or domain name is already in use by another computer on the
	// network.
	ERROR_DUP_DOMAINNAME = 1221

	// The network is not present or not started.
	ERROR_NO_NETWORK = 1222

	// The operation was cancelled by the user.
	ERROR_CANCELLED = 1223

	// The requested operation cannot be performed on a file with a user mapped section open.
	ERROR_USER_MAPPED_FILE = 1224

	// The remote system refused the network connection.
	ERROR_CONNECTION_REFUSED = 1225

	// The network connection was gracefully closed.
	ERROR_GRACEFUL_DISCONNECT = 1226

	// The network transport endpoint already has an address associated with it.
	ERROR_ADDRESS_ALREADY_ASSOCIATED = 1227

	// An address has not yet been associated with the network endpoint.
	ERROR_ADDRESS_NOT_ASSOCIATED = 1228

	// An operation was attempted on a non-existent network connection.
	ERROR_CONNECTION_INVALID = 1229

	// An invalid operation was attempted on an active network connection.
	ERROR_CONNECTION_ACTIVE = 1230

	// The remote network is not reachable by the transport.
	ERROR_NETWORK_UNREACHABLE = 1231

	// The remote system is not reachable by the transport.
	ERROR_HOST_UNREACHABLE = 1232

	// The remote system does not support the transport protocol.
	ERROR_PROTOCOL_UNREACHABLE = 1233

	// No service is operating at the destination network endpoint
	// on the remote system.
	ERROR_PORT_UNREACHABLE = 1234

	// The request was aborted.
	ERROR_REQUEST_ABORTED = 1235

	// The network connection was aborted by the local system.
	ERROR_CONNECTION_ABORTED = 1236

	// The operation could not be completed.  A retry should be performed.
	ERROR_RETRY = 1237

	// A connection to the server could not be made because the limit on the number of
	// concurrent connections for this account has been reached.
	ERROR_CONNECTION_COUNT_LIMIT = 1238

	// Attempting to login during an unauthorized time of day for this account.
	ERROR_LOGIN_TIME_RESTRICTION = 1239

	// The account is not authorized to login from this station.
	ERROR_LOGIN_WKSTA_RESTRICTION = 1240

	// The network address could not be used for the operation requested.
	ERROR_INCORRECT_ADDRESS = 1241

	// The service is already registered.
	ERROR_ALREADY_REGISTERED = 1242

	// The specified service does not exist.
	ERROR_SERVICE_NOT_FOUND = 1243

	// The operation being requested was not performed because the user
	// has not been authenticated.
	ERROR_NOT_AUTHENTICATED = 1244

	// The operation being requested was not performed because the user
	// has not logged on to the network.
	// The specified service does not exist.
	ERROR_NOT_LOGGED_ON = 1245

	// Return that wants caller to continue with work in progress.
	ERROR_CONTINUE = 1246

	// An attempt was made to perform an initialization operation when
	// initialization has already been completed.
	ERROR_ALREADY_INITIALIZED = 1247

	// No more local devices.
	ERROR_NO_MORE_DEVICES = 1248

	// Security Status Codes
	// Not all privileges referenced are assigned to the caller.
	ERROR_NOT_ALL_ASSIGNED = 1300

	// Some mapping between account names and security IDs was not done.
	ERROR_SOME_NOT_MAPPED = 1301

	// No system quota limits are specifically set for this account.
	ERROR_NO_QUOTAS_FOR_ACCOUNT = 1302

	// No encryption key is available.  A well-known encryption key was returned.
	ERROR_LOCAL_USER_SESSION_KEY = 1303

	// The NT password is too complex to be converted to a LAN Manager
	// password.  The LAN Manager password returned is a NULL string.
	ERROR_NULL_LM_PASSWORD = 1304

	// The revision level is unknown.
	ERROR_UNKNOWN_REVISION = 1305

	// Indicates two revision levels are incompatible.
	ERROR_REVISION_MISMATCH = 1306

	// This security ID may not be assigned as the owner of this object.
	ERROR_INVALID_OWNER = 1307

	// This security ID may not be assigned as the primary group of an object.
	ERROR_INVALID_PRIMARY_GROUP = 1308

	// An attempt has been made to operate on an impersonation token
	// by a thread that is not currently impersonating a client.
	ERROR_NO_IMPERSONATION_TOKEN = 1309

	// The group may not be disabled.
	ERROR_CANT_DISABLE_MANDATORY = 1310

	// There are currently no logon servers available to service the logon
	// request.
	ERROR_NO_LOGON_SERVERS = 1311

	//  A specified logon session does not exist.  It may already have
	//  been terminated.
	ERROR_NO_SUCH_LOGON_SESSION = 1312

	//  A specified privilege does not exist.
	ERROR_NO_SUCH_PRIVILEGE = 1313

	//  A required privilege is not held by the client.
	ERROR_PRIVILEGE_NOT_HELD = 1314

	// The name provided is not a properly formed account name.
	ERROR_INVALID_ACCOUNT_NAME = 1315

	// The specified user already exists.
	ERROR_USER_EXISTS = 1316

	// The specified user does not exist.
	ERROR_NO_SUCH_USER = 1317

	// The specified group already exists.
	ERROR_GROUP_EXISTS = 1318

	// The specified group does not exist.
	ERROR_NO_SUCH_GROUP = 1319

	// Either the specified user account is already a member of the specified
	// group, or the specified group cannot be deleted because it contains
	// a member.
	ERROR_MEMBER_IN_GROUP = 1320

	// The specified user account is not a member of the specified group account.
	ERROR_MEMBER_NOT_IN_GROUP = 1321

	// The last remaining administration account cannot be disabled
	// or deleted.
	ERROR_LAST_ADMIN = 1322

	// Unable to update the password.  The value provided as the current
	// password is incorrect.
	ERROR_WRONG_PASSWORD = 1323

	// Unable to update the password.  The value provided for the new password
	// contains values that are not allowed in passwords.
	ERROR_ILL_FORMED_PASSWORD = 1324

	// Unable to update the password because a password update rule has been
	// violated.
	ERROR_PASSWORD_RESTRICTION = 1325

	// Logon failure: unknown user name or bad password.
	ERROR_LOGON_FAILURE = 1326

	// Logon failure: user account restriction.
	ERROR_ACCOUNT_RESTRICTION = 1327

	// Logon failure: account logon time restriction violation.
	ERROR_INVALID_LOGON_HOURS = 1328

	// Logon failure: user not allowed to log on to this computer.
	ERROR_INVALID_WORKSTATION = 1329

	// Logon failure: the specified account password has expired.
	ERROR_PASSWORD_EXPIRED = 1330

	// Logon failure: account currently disabled.
	ERROR_ACCOUNT_DISABLED = 1331

	// No mapping between account names and security IDs was done.
	ERROR_NONE_MAPPED = 1332

	// Too many local user identifiers (LUIDs) were requested at one time.
	ERROR_TOO_MANY_LUIDS_REQUESTED = 1333

	// No more local user identifiers (LUIDs) are available.
	ERROR_LUIDS_EXHAUSTED = 1334

	// The subauthority part of a security ID is invalid for this particular use.
	ERROR_INVALID_SUB_AUTHORITY = 1335

	// The access control list (ACL) structure is invalid.
	ERROR_INVALID_ACL = 1336

	// The security ID structure is invalid.
	ERROR_INVALID_SID = 1337

	// The security descriptor structure is invalid.
	ERROR_INVALID_SECURITY_DESCR = 1338

	// The inherited access control list (ACL) or access control entry (ACE)
	// could not be built.
	ERROR_BAD_INHERITANCE_ACL = 1340

	// The server is currently disabled.
	ERROR_SERVER_DISABLED = 1341

	// The server is currently enabled.
	ERROR_SERVER_NOT_DISABLED = 1342

	// The value provided was an invalid value for an identifier authority.
	ERROR_INVALID_ID_AUTHORITY = 1343

	// No more memory is available for security information updates.
	ERROR_ALLOTTED_SPACE_EXCEEDED = 1344

	// The specified attributes are invalid, or incompatible with the
	// attributes for the group as a whole.
	ERROR_INVALID_GROUP_ATTRIBUTES = 1345

	// Either a required impersonation level was not provided, or the
	// provided impersonation level is invalid.
	ERROR_BAD_IMPERSONATION_LEVEL = 1346

	// Cannot open an anonymous level security token.
	ERROR_CANT_OPEN_ANONYMOUS = 1347

	// The validation information class requested was invalid.
	ERROR_BAD_VALIDATION_CLASS = 1348

	// The type of the token is inappropriate for its attempted use.
	ERROR_BAD_TOKEN_TYPE = 1349

	// Unable to perform a security operation on an object
	// which has no associated security.
	ERROR_NO_SECURITY_ON_OBJECT = 1350

	// Indicates a Windows NT Server could not be contacted or that
	// objects within the domain are protected such that necessary
	// information could not be retrieved.
	ERROR_CANT_ACCESS_DOMAIN_INFO = 1351

	// The security account manager (SAM) or local security
	// authority (LSA) server was in the wrong state to perform
	// the security operation.
	ERROR_INVALID_SERVER_STATE = 1352

	// The domain was in the wrong state to perform the security operation.
	ERROR_INVALID_DOMAIN_STATE = 1353

	// This operation is only allowed for the Primary Domain Controller of the domain.
	ERROR_INVALID_DOMAIN_ROLE = 1354

	// The specified domain did not exist.
	ERROR_NO_SUCH_DOMAIN = 1355

	// The specified domain already exists.
	ERROR_DOMAIN_EXISTS = 1356

	// An attempt was made to exceed the limit on the number of domains per server.
	ERROR_DOMAIN_LIMIT_EXCEEDED = 1357

	// Unable to complete the requested operation because of either a
	// catastrophic media failure or a data structure corruption on the disk.
	ERROR_INTERNAL_DB_CORRUPTION = 1358

	// The security account database contains an internal inconsistency.
	ERROR_INTERNAL_ERROR = 1359

	// Generic access types were contained in an access mask which should
	// already be mapped to non-generic types.
	ERROR_GENERIC_NOT_MAPPED = 1360

	// A security descriptor is not in the right format (absolute or self-relative).
	ERROR_BAD_DESCRIPTOR_FORMAT = 1361

	// The requested action is restricted for use by logon processes
	// only.  The calling process has not registered as a logon process.
	ERROR_NOT_LOGON_PROCESS = 1362

	// Cannot start a new logon session with an ID that is already in use.
	ERROR_LOGON_SESSION_EXISTS = 1363

	// A specified authentication package is unknown.
	ERROR_NO_SUCH_PACKAGE = 1364

	// The logon session is not in a state that is consistent with the
	// requested operation.
	ERROR_BAD_LOGON_SESSION_STATE = 1365

	// The logon session ID is already in use.
	ERROR_LOGON_SESSION_COLLISION = 1366

	// A logon request contained an invalid logon type value.
	ERROR_INVALID_LOGON_TYPE = 1367

	// Unable to impersonate via a named pipe until data has been read
	// from that pipe.
	ERROR_CANNOT_IMPERSONATE = 1368

	// The transaction state of a Registry subtree is incompatible with the
	// requested operation.
	ERROR_RXACT_INVALID_STATE = 1369

	// An internal security database corruption has been encountered.
	ERROR_RXACT_COMMIT_FAILURE = 1370

	// Cannot perform this operation on built-in accounts.
	ERROR_SPECIAL_ACCOUNT = 1371

	// Cannot perform this operation on this built-in special group.
	ERROR_SPECIAL_GROUP = 1372

	// Cannot perform this operation on this built-in special user.
	ERROR_SPECIAL_USER = 1373

	// The user cannot be removed from a group because the group
	// is currently the user's primary group.
	ERROR_MEMBERS_PRIMARY_GROUP = 1374

	// The token is already in use as a primary token.
	ERROR_TOKEN_ALREADY_IN_USE = 1375

	// The specified local group does not exist.
	ERROR_NO_SUCH_ALIAS = 1376

	// The specified account name is not a member of the local group.
	ERROR_MEMBER_NOT_IN_ALIAS = 1377

	// The specified account name is already a member of the local group.
	ERROR_MEMBER_IN_ALIAS = 1378

	// The specified local group already exists.
	ERROR_ALIAS_EXISTS = 1379

	// Logon failure: the user has not been granted the requested
	// logon type at this computer.
	ERROR_LOGON_NOT_GRANTED = 1380

	// The maximum number of secrets that may be stored in a single system has been
	// exceeded.
	ERROR_TOO_MANY_SECRETS = 1381

	// The length of a secret exceeds the maximum length allowed.
	ERROR_SECRET_TOO_LONG = 1382

	// The local security authority database contains an internal inconsistency.
	ERROR_INTERNAL_DB_ERROR = 1383

	// During a logon attempt, the user's security context accumulated too many
	// security IDs.
	ERROR_TOO_MANY_CONTEXT_IDS = 1384

	// Logon failure: the user has not been granted the requested logon type
	// at this computer.
	ERROR_LOGON_TYPE_NOT_GRANTED = 1385

	// A cross-encrypted password is necessary to change a user password.
	ERROR_NT_CROSS_ENCRYPTION_REQUIRED = 1386

	// A new member could not be added to a local group because the member does
	// not exist.
	ERROR_NO_SUCH_MEMBER = 1387

	// A new member could not be added to a local group because the member has the
	// wrong account type.
	ERROR_INVALID_MEMBER = 1388

	// Too many security IDs have been specified.
	ERROR_TOO_MANY_SIDS = 1389

	// A cross-encrypted password is necessary to change this user password.
	ERROR_LM_CROSS_ENCRYPTION_REQUIRED = 1390

	// Indicates an TACL contains no inheritable components
	ERROR_NO_INHERITANCE = 1391

	// The file or directory is corrupt and non-readable.
	ERROR_FILE_CORRUPT = 1392

	// The disk structure is corrupt and non-readable.
	ERROR_DISK_CORRUPT = 1393

	// There is no user session key for the specified logon session.
	ERROR_NO_USER_SESSION_KEY = 1394

	// The service being accessed is licensed for a particular number of connections.
	// No more connections can be made to the service at this time
	// because there are already as many connections as the service can accept.
	ERROR_LICENSE_QUOTA_EXCEEDED = 1395

	// WinUser Error Codes
	// Invalid window handle.
	ERROR_INVALID_WINDOW_HANDLE = 1400

	// Invalid menu handle.
	ERROR_INVALID_MENU_HANDLE = 1401

	// Invalid cursor handle.
	ERROR_INVALID_CURSOR_HANDLE = 1402

	// Invalid accelerator table handle.
	ERROR_INVALID_ACCEL_HANDLE = 1403

	// Invalid hook handle.
	ERROR_INVALID_HOOK_HANDLE = 1404

	// Invalid handle to a multiple-window position structure.
	ERROR_INVALID_DWP_HANDLE = 1405

	// Cannot create a top-level child window.
	ERROR_TLW_WITH_WSCHILD = 1406

	// Cannot find window class.
	ERROR_CANNOT_FIND_WND_CLASS = 1407

	// Invalid window, belongs to other thread.
	ERROR_WINDOW_OF_OTHER_THREAD = 1408

	// Hot key is already registered.
	ERROR_HOTKEY_ALREADY_REGISTERED = 1409

	// Class already exists.
	ERROR_CLASS_ALREADY_EXISTS = 1410

	// Class does not exist.
	ERROR_CLASS_DOES_NOT_EXIST = 1411

	// Class still has open windows.
	ERROR_CLASS_HAS_WINDOWS = 1412

	// Invalid index.
	ERROR_INVALID_INDEX = 1413

	// Invalid icon handle.
	ERROR_INVALID_ICON_HANDLE = 1414

	// Using private DIALOG window words.
	ERROR_PRIVATE_DIALOG_INDEX = 1415

	// The listbox identifier was not found.
	ERROR_LISTBOX_ID_NOT_FOUND = 1416

	// No wildcards were found.
	ERROR_NO_WILDCARD_CHARACTERS = 1417

	// Thread does not have a clipboard open.
	ERROR_CLIPBOARD_NOT_OPEN = 1418

	// Hot key is not registered.
	ERROR_HOTKEY_NOT_REGISTERED = 1419

	// The window is not a valid dialog window.
	ERROR_WINDOW_NOT_DIALOG = 1420

	// Control ID not found.
	ERROR_CONTROL_ID_NOT_FOUND = 1421

	// Invalid message for a combo box because it does not have an edit control.
	ERROR_INVALID_COMBOBOX_MESSAGE = 1422

	// The window is not a combo box.
	ERROR_WINDOW_NOT_COMBOBOX = 1423

	// Height must be less than 256.
	ERROR_INVALID_EDIT_HEIGHT = 1424

	// Invalid device context (DC) handle.
	ERROR_DC_NOT_FOUND = 1425

	// Invalid hook procedure type.
	ERROR_INVALID_HOOK_FILTER = 1426

	// Invalid hook procedure.
	ERROR_INVALID_FILTER_PROC = 1427

	// Cannot set non-local hook without a module handle.
	ERROR_HOOK_NEEDS_HMOD = 1428

	// This hook procedure can only be set globally.
	ERROR_GLOBAL_ONLY_HOOK = 1429

	// The journal hook procedure is already installed.
	ERROR_JOURNAL_HOOK_SET = 1430

	// The hook procedure is not installed.
	ERROR_HOOK_NOT_INSTALLED = 1431

	// Invalid message for single-selection listbox.
	ERROR_INVALID_LB_MESSAGE = 1432

	// LB_SETCOUNT sent to non-lazy listbox.
	ERROR_SETCOUNT_ON_BAD_LB = 1433

	// This list box does not support tab stops.
	ERROR_LB_WITHOUT_TABSTOPS = 1434

	// Cannot destroy object created by another thread.
	ERROR_DESTROY_OBJECT_OF_OTHER_THREAD = 1435

	// Child windows cannot have menus.
	ERROR_CHILD_WINDOW_MENU = 1436

	// The window does not have a system menu.
	ERROR_NO_SYSTEM_MENU = 1437

	// Invalid message box style.
	ERROR_INVALID_MSGBOX_STYLE = 1438

	// Invalid system-wide (SPI_*) parameter.
	ERROR_INVALID_SPI_VALUE = 1439

	// Screen already locked.
	ERROR_SCREEN_ALREADY_LOCKED = 1440

	// All handles to windows in a multiple-window position structure must
	// have the same parent.
	ERROR_HWNDS_HAVE_DIFF_PARENT = 1441

	// The window is not a child window.
	ERROR_NOT_CHILD_WINDOW = 1442

	// Invalid GW_* command.
	ERROR_INVALID_GW_COMMAND = 1443

	// Invalid thread identifier.
	ERROR_INVALID_THREAD_ID = 1444

	// Cannot process a message from a window that is not a multiple document
	// interface (MDI) window.
	ERROR_NON_MDICHILD_WINDOW = 1445

	// Popup menu already active.
	ERROR_POPUP_ALREADY_ACTIVE = 1446

	// The window does not have scroll bars.
	ERROR_NO_SCROLLBARS = 1447

	// Scroll bar range cannot be greater than $7FFF.
	ERROR_INVALID_SCROLLBAR_RANGE = 1448

	// Cannot show or remove the window in the way specified.
	ERROR_INVALID_SHOWWIN_COMMAND = 1449

	// Insufficient system resources exist to complete the requested service.
	ERROR_NO_SYSTEM_RESOURCES = 1450

	// Insufficient system resources exist to complete the requested service.
	ERROR_NONPAGED_SYSTEM_RESOURCES = 1451

	// Insufficient system resources exist to complete the requested service.
	ERROR_PAGED_SYSTEM_RESOURCES = 1452

	// Insufficient quota to complete the requested service.
	ERROR_WORKING_SET_QUOTA = 1453

	// Insufficient quota to complete the requested service.
	ERROR_PAGEFILE_QUOTA = 1454

	// The paging file is too small for this operation to complete.
	ERROR_COMMITMENT_LIMIT = 1455

	// A menu item was not found.
	ERROR_MENU_ITEM_NOT_FOUND = 1456

	// Invalid keyboard layout handle.
	ERROR_INVALID_KEYBOARD_HANDLE = 1457

	// Hook type not allowed.
	ERROR_HOOK_TYPE_NOT_ALLOWED = 1458

	// This operation requires an interactive windowstation.
	ERROR_REQUIRES_INTERACTIVE_WINDOWSTATION = 1459

	// This operation returned because the timeout period expired.
	ERROR_TIMEOUT = 1460

	// Eventlog Status Codes
	// The event log file is corrupt.
	ERROR_EVENTLOG_FILE_CORRUPT = 1500

	// No event log file could be opened, so the event logging service did not start.
	ERROR_EVENTLOG_CANT_START = 1501

	// The event log file is full.
	ERROR_LOG_FILE_FULL = 1502

	// The event log file has changed between reads.
	ERROR_EVENTLOG_FILE_CHANGED = 1503
)
